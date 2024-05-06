package mock_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-github/v61/github"
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

		// The rate limiter will allow 10 requests per second, and a burst size of 1.
		// These two options together mean that the rate of requests will be strictly enforced, so if any two requests are
		// made less than 1/10th of a second apart, the latter will be refused and come back with a rate limit error.
		mock.WithRateLimit(10, 1),
	)

	ghc := github.NewClient(mhc)
	opts := &github.RepositoryListByOrgOptions{}
	repoNames := []string{}
	rleCount := 0

	for {
		repos, resp, err := ghc.Repositories.ListByOrg(context.Background(), "org", opts)
		if err != nil {
			// The only type of error we expect is one from the rate limiter.
			var rle *github.RateLimitError
			if !errors.As(err, &rle) {
				t.Fatalf("error is not a github.RateLimitError: %v", err)
			}

			rleCount++

			// After hitting the rate limiter and getting the appropriate error above, sleeping for one tenth of a second
			// should be enough to reset the limiter and try requesting the repo list again.
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
