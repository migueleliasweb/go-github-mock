# go-github-mock
[![Go Reference](https://pkg.go.dev/badge/github.com/migueleliasweb/go-github-mock.svg)](https://pkg.go.dev/github.com/migueleliasweb/go-github-mock) [![Go Report Card](https://goreportcard.com/badge/github.com/migueleliasweb/go-github-mock)](https://goreportcard.com/report/github.com/migueleliasweb/go-github-mock)

A library to aid unittesting code that uses Golang's Github SDK

# Installation

```bash
go get github.com/migueleliasweb/go-github-mock
```

# Features

- Create mocks for successive calls for the same endpoint
- Pagination support
- Mock error returns
- High level abstraction helps writing readabe unittests (see `mock.WithRequestMatch`)
- Lower level abstraction for advanced uses (see `mock.WithRequestMatchHandler`)

# Breaking changes

- `v0.0.3` the API for the server options have beem simplified
- `v0.0.4` fixes to the gen script caused multiple url matches to change

# Example

```
import "github.com/migueleliasweb/go-github-mock/src/mock"
```

## Multiple requests

```golang
mockedHTTPClient := mock.NewMockedHTTPClient(
    mock.WithRequestMatch(
        mock.GetUsersByUsername,
        github.User{
            Name: github.String("foobar"),
        },
    ),
    mock.WithRequestMatch(
        mock.GetUsersOrgsByUsername,
        []github.Organization{
            {
                Name: github.String("foobar123thisorgwasmocked"),
            },
        },
    ),
    mock.WithRequestMatchHandler(
        mock.GetOrgsProjectsByOrg,
        http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
            w.Write(mock.MustMarshal([]github.Project{
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

user, _, userErr := c.Users.Get(ctx, "myuser")

// user.Name == "foobar"

orgs, _, orgsErr := c.Organizations.List(
    ctx,
    *(user.Name),
    nil,
)

// orgs[0].Name == "foobar123thisorgwasmocked"

projs, _, projsErr := c.Organizations.ListProjects(
    ctx,
    *orgs[0].Name,
    &github.ProjectListOptions{},
)

// projs[0].Name == "mocked-proj-1"
// projs[1].Name == "mocked-proj-2"

```

## Returning empty results

```golang
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

// len(issues1) == 2
// repo1Err == nil

issues2, _, repo2Err := c.Issues.ListByRepo(ctx, "owner1", "repo2", &github.IssueListByRepoOptions{})

// len(issues2) == 0
// repo2Err == nil
```

## Mocking errors from the API

```golang
mockedHTTPClient := mock.NewMockedHTTPClient(
    mock.WithRequestMatchHandler(
        mock.GetUsersByUsername,
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            mock.WriteError(
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

// user == nil

if userErr == nil {	
    if ghErr, ok := userErr.(*github.ErrorResponse); ok {
        fmt.Println(ghErr.Message) // == "github went belly up or something"
    }
}

```

## Mocking with pagination

```golang
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
        // in fact, the perPage option is ignored my the mocks
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

    // len(repos) == 2

    allRepos = append(allRepos, repos...)

    if resp.NextPage == 0 {
        break
    }

    opt.Page = resp.NextPage
}

// matches the mock definitions len(page[0]) + len(page[1])
// len(allRepos) == 4
```

## Mocking for Github Enterprise

Github Enterprise uses a different prefix for its endpoints. In order to use the correct endpoints, please use the different set of `*Enterprise` options:

- WithRequestMatchEnterprise
- WithRequestMatchPagesEnterprise

```golang
mockedHTTPClient := mock.NewMockedHTTPClient(
    mock.WithRequestMatchEnterprise( // uses enterprise endpoints instead
        mock.GetUsersByUsername,
        github.User{
            Name: github.String("foobar"),
        },
    ),
)

c := githubEnterprise.NewClient(mockedHTTPClient)

ctx := context.Background()

user, _, userErr := c.Users.Get(ctx, "myuser")

// user.Name == "foobar"
```

# Why

Some conversations got started on [go-github#1800](https://github.com/google/go-github/issues/1800) since `go-github` didn't provide an interface that could be easily reimplemented for unittests. After lots of conversations from the folks from [go-github](https://github.com/google/go-github) and quite a few PR ideas later, this style of testing was deemed not suitable to be part of the core SDK as it's not a feature of the API itself. Nonetheless, the ability of writing unittests for code that uses the `go-github` package is critical. 

A reuseable, and not overly verbose, way of writing the tests was reached after some more interactions (months down the line) and here we are.

# Thanks

Thanks for all ideas and feedback from the folks in [go-github](https://github.com/google/go-github/).

# License

This library is distributed under the MIT License found in LICENSE.
