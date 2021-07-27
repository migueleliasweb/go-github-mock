# go-github-mock
A library to aid unittesting code that uses Golang's Github SDK

# Installation

```bash
go get github.com/migueleliasweb/go-github-mock
```

# Features

- Create mocks for successive calls for the same endpoint
- Mock error returns
- High level abstraction helps writing readabe unittests (see `mock.WithRequestMatch`)
- Lower level abstraction for advanced uses (see `mock.WithRequestMatchHandler`)

# Example

```
import "github.com/migueleliasweb/go-github-mock/src/mock"
```

## Multiple requests

```golang
mockedHttpClient := mock.NewMockedHttpClient(
    mock.WithRequestMatch(
        mock.GetUsersByUsername,
        [][]byte{
            mock.MustMarshal(github.User{
                Name: github.String("foobar"),
            }),
        },
    ),
    mock.WithRequestMatch(
        mock.GetUsersOrgsByUsername,
        [][]byte{
            mock.MustMarshal([]github.Organization{
                {
                    Name: github.String("foobar123thisorgwasmocked"),
                },
            }),
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
c := github.NewClient(mockedHttpClient)

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

## Mocking errors from the API

```golang
mockedHttpClient := mock.NewMockedHttpClient(
    mock.WithRequestMatchHandler(
        mock.GetUsersByUsername,
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            WriteError(
                w,
                http.StatusInternalServerError,
                "github went belly up or something",
            )
        }),
    ),
)
c := github.NewClient(mockedHttpClient)

ctx := context.Background()

user, _, userErr := c.Users.Get(ctx, "someUser")

// user == nil

if userErr == nil {	
    if ghErr, ok := userErr.(*github.ErrorResponse); ok {
        fmt.Println(ghErr.Message) // == "github went belly up or something"
    }
}

```

# Why

Some conversations got started on [go-github#1800](https://github.com/google/go-github/issues/1800) since `go-github` didn't provide an interface that could be easily reimplemented for unittests. After lots of conversations from the folks from [go-github](https://github.com/google/go-github) and quite a few PR ideas later, this style of testing was deemed not suitable to be part of the core SDK as it's not a feature of the API itself. Nonetheless, the ability of writing unittests for code that uses the `go-github` package is critical. 

A reuseable, and not overly verbose, way of writing the tests was reached after some more conversations (months down the line) and here we are.

# Thanks

Thanks for all ideas and feedback from the folks in [go-github](https://github.com/google/go-github/).

# License

This library is distributed under the MIT License found in LICENSE.