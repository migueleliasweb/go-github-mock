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
				MustMarshal(github.Organization{
					Name: github.String("foobar123thisorgwasmocked"),
				}),
			},
		),
	)
	c := github.NewClient(mockedHttpClient)

	ctx := context.Background()

	user, _, userErr := c.Users.Get(ctx, "someUser")

	if user == nil || user.Name == nil || *user.Name != "foobar" {
		t.Fatalf("User name is %s, want foobar", user)
	}

	if userErr != nil {
		t.Errorf("User err is %s, want nil", userErr.Error())
	}

	orgs, _, orgsErr := c.Organizations.List(
		ctx,
		*(user.Name),
		nil,
	)

	if len(orgs) != 1 {
		t.Errorf("Orgs len is %d want 1", len(orgs))
	}

	if orgsErr != nil {
		t.Errorf("Err is %s, want nil", orgsErr.Error())
	}

	if *(orgs[0].Name) != "foobar123thisorgwasmocked" {
		t.Errorf("orgs[0].Name is %s, want %s", *orgs[0].Name, "foobar123thisorgdoesnotexist")
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
