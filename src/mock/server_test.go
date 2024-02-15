package mock

import (
	"context"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-github/v59/github"
)

func TestNewMockedHTTPClient(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatch(
			GetUsersByUsername,
			github.User{
				Name: github.String("foobar"),
			},
		),
		WithRequestMatch(
			GetUsersOrgsByUsername,
			[]github.Organization{
				{
					Name: github.String("foobar123thisorgwasmocked"),
				},
			},
		),
		WithRequestMatchHandler(
			GetOrgsProjectsByOrg,
			http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Write(MustMarshal([]github.Project{
					{
						Name: github.String("mocked-proj-1"),
					},
					{
						Name: github.String("mocked-proj-2"),
					},
				}))
			}),
		),
	)
	c := github.NewClient(mockedHTTPClient)

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

func TestMockErrorSimple(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatchHandler(
			GetUsersByUsername,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				WriteError(
					w,
					http.StatusInternalServerError,
					"github went belly up or something",
				)
			}),
		),
	)
	c := github.NewClient(mockedHTTPClient)

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

func TestMockErrorToError(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatchHandler(
			GetUsersByUsername,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				WriteError(
					w,
					http.StatusInternalServerError,
					"github went belly up or something",
				)
			}),
		),
	)
	c := github.NewClient(mockedHTTPClient)

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

	// We should be able to use the default formatting for the error interface.
	// This was initially failing due to the `github.ErrorResponse` containing `Response: nil`
	// thus making the folling call panic due to `nil pointer`.
	if fmtdErr := userErr.Error(); len(fmtdErr) == 0 {
		t.Fatal("couldn't use default error formattung for github.ErrorResponse")
	}
}

func TestMocksNotConfiguredError(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatch(
			GetUsersByUsername,
			github.User{
				Name: github.String("foobar"),
			},
		),
	)
	c := github.NewClient(mockedHTTPClient)

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

func TestMocksPaginationAllPages(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatchPages(
			GetOrgsReposByOrg,
			[]github.Repository{
				{
					Name: github.String("repo-A-on-first-page"),
				},
				{
					Name: github.String("repo-B-on-first-page"),
				},
			},
			[]github.Repository{
				{
					Name: github.String("repo-C-on-second-page"),
				},
				{
					Name: github.String("repo-D-on-second-page"),
				},
			},
		),
	)

	c := github.NewClient(mockedHTTPClient)

	ctx := context.Background()

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{
			// in fact, the perPage option is ignored
			// but this would be present in production code
			PerPage: 2,
		},
	}

	var allRepos []*github.Repository

	for {
		repos, resp, listErr := c.Repositories.ListByOrg(ctx, "foobar", opt)

		if listErr != nil {
			t.Errorf("error listing repositories: %s", listErr.Error())
		}

		// matches mock definition
		if len(repos) != 2 {
			t.Errorf("len(repos) is %d, want 2", len(repos))
		}

		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	if len(allRepos) != 4 {
		t.Errorf("len(allRepos) is %d, want 4", len(allRepos))
	}
}

func TestEmptyArrayResult(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatch(
			GetReposIssuesByOwnerByRepo,
			[]github.Issue{
				{
					ID:    github.Int64(123),
					Title: github.String("Issue 1"),
				},
				{
					ID:    github.Int64(456),
					Title: github.String("Issue 2"),
				},
			},
			[]github.Issue{},
		),
	)

	c := github.NewClient(mockedHTTPClient)

	ctx := context.Background()

	issues1, _, repo1Err := c.Issues.ListByRepo(ctx, "owner1", "repo1", &github.IssueListByRepoOptions{})

	if repo1Err != nil {
		t.Errorf("error listing repository1 issues: %s", repo1Err.Error())
	}

	if len(issues1) != 2 {
		t.Errorf("len(issues1) is %d, want 2", len(issues1))
	}

	issues2, _, repo2Err := c.Issues.ListByRepo(ctx, "owner1", "repo2", &github.IssueListByRepoOptions{})

	if repo2Err != nil {
		t.Errorf("error listing repository2 issues: %s", repo2Err.Error())
	}

	if len(issues2) != 0 {
		t.Errorf("len(issues2) is %d, want 0", len(issues2))
	}
}
