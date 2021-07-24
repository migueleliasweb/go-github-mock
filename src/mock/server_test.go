package mock

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/v37/github"
)

func TestNewMockedHttpClient(t *testing.T) {
	mockedHttpClient := NewMockedHttpClient(
		WithRequestMatch(
			GetUsersByUsername,
			[][]byte{
				MustMarshal(github.User{
					Name: github.String("foobar"),
				}),
			},
		),
		WithRequestMatch(
			GetUsersOrgsByUsername,
			[][]byte{
				MustMarshal([]github.Organization{
					{
						Name: github.String("foobar123thisorgwasmocked"),
					},
				}),
			},
		),
		WithRequestMatchHandler(
			GetOrgsProjectsByOrg,
			func(w http.ResponseWriter, _ *http.Request) {
				w.Write(MustMarshal([]github.Project{
					{
						Name: github.String("mocked-proj-1"),
					},
					{
						Name: github.String("mocked-proj-2"),
					},
				}))
			},
		),
	)
	c := github.NewClient(mockedHttpClient)

	ctx := context.Background()

	user, _, userErr := c.Users.Get(ctx, "someUser")

	if user == nil || user.Name == nil || *user.Name != "foobar" {
		t.Fatalf("user name is %s, want foobar", user)
	}

	if userErr != nil {
		t.Errorf("user err is %s, want nil", userErr.Error())
	}

	orgs, _, orgsErr := c.Organizations.List(
		ctx,
		*(user.Name),
		nil,
	)

	if len(orgs) != 1 {
		t.Errorf("orgs len is %d want 1", len(orgs))
	}

	if orgsErr != nil {
		t.Errorf("orgs err is %s, want nil", orgsErr.Error())
	}

	if *(orgs[0].Name) != "foobar123thisorgwasmocked" {
		t.Errorf("orgs[0].Name is %s, want %s", *orgs[0].Name, "foobar123thisorgdoesnotexist")
	}

	projs, _, projsErr := c.Organizations.ListProjects(
		ctx,
		*orgs[0].Name,
		&github.ProjectListOptions{},
	)

	if projsErr != nil {
		t.Errorf("projs err is %s, want nil", projsErr.Error())
	}

	if len(projs) != 2 {
		t.Errorf("projs len is %d want 2", len(projs))
	}
}

func TestExampleServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	c := ts.Client()

	c.Transport = &EnforceHostRoundTripper{
		Host:                 ts.URL,
		UpstreamRoundTripper: ts.Client().Transport,
	}

	res, err := c.Get("https://docs.github.com")

	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}
