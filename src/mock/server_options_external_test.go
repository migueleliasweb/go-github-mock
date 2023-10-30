package mock_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-github/v56/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
)

func TestWithRateLimit(t *testing.T) {
	t.Parallel()

	const (
		repoOne = "repo-one"
		repoTwo = "repo-two"
	)

	mhc := mock.NewMockedHTTPClient(
		mock.WithRequestMatchPages(
			mock.GetOrgsReposByOrg,
			[]github.Repository{{Name: github.String(repoOne)}},
			[]github.Repository{{Name: github.String(repoTwo)}},
		),

		// The rate limiter will allow 10 requests per second.
		mock.WithRateLimit(10),
	)

	ghc := github.NewClient(mhc)
	opts := &github.RepositoryListByOrgOptions{}
	repoNames := []string{}
	rleCount := 0

	for {
		repos, resp, err := ghc.Repositories.ListByOrg(context.Background(), "org", opts)
		if err != nil {
			// The only type of error we expect is one from the rate limiter.
			if _, ok := err.(*github.RateLimitError); !ok {
				t.Fatalf("error is not a github.RateLimitError: %v", err)
			}

			rleCount++

			// Sleeping for one tenth of a second should be enough for the rate limiter on the mock server.
			time.Sleep(time.Second / 10)

			continue
		}
		defer resp.Body.Close()

		if len(repos) != 1 {
			t.Fatal("should receive one repo per page from mock")
		}

		// Save the returned repo name to our slice, to assert against later.
		repoNames = append(repoNames, repos[0].GetName())

		// Handle pagination.
		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	if len(repoNames) != 2 || repoNames[0] != repoOne || repoNames[1] != repoTwo {
		t.Fatalf("list of returned repo names is wrong: %v", repoNames)
	}

	if rleCount != 1 {
		t.Fatal("did not receive the expected number of github.RateLimitError")
	}
}
