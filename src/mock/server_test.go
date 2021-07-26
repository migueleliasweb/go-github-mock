package mock

import (
	"context"
	"net/http"
	"reflect"
	"strings"
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

func TestMockErrors(t *testing.T) {
	mockedHttpClient := NewMockedHttpClient(
		WithRequestMatchHandler(
			GetUsersByUsername,
			func(w http.ResponseWriter, r *http.Request) {
				WriteError(
					w,
					http.StatusInternalServerError,
					"github went belly up or something",
				)
			},
		),
	)
	c := github.NewClient(mockedHttpClient)

	ctx := context.Background()

	user, _, userErr := c.Users.Get(ctx, "someUser")

	if userV := reflect.ValueOf(user); !userV.IsZero() {
		t.Errorf("user is %s, want nil", user)
	}

	if userErr == nil {
		t.Fatal("user err is nil, want *github.ErrorResponse")
	}

	ghErr, ok := userErr.(*github.ErrorResponse)

	if !ok {
		t.Fatal("couldn't cast userErr to *github.ErrorResponse")
	}

	if ghErr.Message != "github went belly up or something" {
		t.Errorf("user err is %s, want 'github went belly up or something'", userErr.Error())
	}
}

func TestMocksNotConfiguredError(t *testing.T) {
	mockedHttpClient := NewMockedHttpClient(
		WithRequestMatch(
			GetUsersByUsername,
			[][]byte{
				MustMarshal(github.User{
					Name: github.String("foobar"),
				}),
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
		"someuser",
		&github.ListOptions{},
	)

	if len(orgs) > 0 {
		t.Errorf("orgs len is %d, want 0", len(orgs))
	}

	if r, ok := orgsErr.(*github.ErrorResponse); ok {
		if !strings.Contains(r.Message, "mock response not found for") {
			t.Errorf("error message should contain 'mock response not found for'")
		}
	}
}
