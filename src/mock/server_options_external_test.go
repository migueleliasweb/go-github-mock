package mock_test

import (
	"context"
	"errors"
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
				t.Fatalf("error from listing repos is not a github.RateLimitError: %v", err)
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

func TestWithRequestMatchOptionsPages(t *testing.T) {
	t.Parallel()

	const (
		repoOne = "repo-one"
		repoTwo = "repo-two"
		teamOne = "team-one"
		teamTwo = "team-two"
	)

	mhc := mock.NewMockedHTTPClient(
		// The rate limiter will allow 10 requests per second to the teams endpoint only.
		mock.WithRequestMatchOptionsPages(
			mock.GetOrgsTeamsByOrg,
			mock.RequestMatchOptions{RPS: 10, Burst: 1},
			[]github.Team{{Name: github.String(teamOne)}},
			[]github.Team{{Name: github.String(teamTwo)}},
		),

		// The repos endpoint should remain unlimited.
		mock.WithRequestMatchPages(
			mock.GetOrgsReposByOrg,
			[]github.Repository{{Name: github.String(repoOne)}},
			[]github.Repository{{Name: github.String(repoTwo)}},
		),
	)

	ghc := github.NewClient(mhc)
	repoOpts := &github.RepositoryListByOrgOptions{}
	teamOpts := &github.ListOptions{}
	repoNames, teamNames := []string{}, []string{}
	rleCount := 0

	for {
		repos, resp, err := ghc.Repositories.ListByOrg(context.Background(), "org", repoOpts)
		if err != nil {
			t.Fatalf("should not error while listing repos: %v", err)
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

		repoOpts.Page = resp.NextPage
	}

	if len(repoNames) != 2 || repoNames[0] != repoOne || repoNames[1] != repoTwo {
		t.Fatalf("list of returned repo names is wrong: %v", repoNames)
	}

	for {
		teams, resp, err := ghc.Teams.ListTeams(context.Background(), "org", teamOpts)
		if err != nil {
			// The only type of error we expect is one from the rate limiter.
			var rle *github.RateLimitError
			if !errors.As(err, &rle) {
				t.Fatalf("error from listing teams is not a github.RateLimitError: %v", err)
			}

			rleCount++

			// After hitting the rate limiter and getting the appropriate error above, sleeping for one tenth of a second
			// should be enough to reset the limiter and try requesting the repo list again.
			time.Sleep(time.Second / 10)

			continue
		}
		defer resp.Body.Close()

		if len(teams) != 1 {
			t.Fatal("should receive one team per page from mock")
		}

		// Save the returned team name to our slice, to assert against later.
		teamNames = append(teamNames, teams[0].GetName())

		// Handle pagination.
		if resp.NextPage == 0 {
			break
		}

		teamOpts.Page = resp.NextPage
	}

	if len(teamNames) != 2 || teamNames[0] != teamOne || teamNames[1] != teamTwo {
		t.Fatalf("list of returned team names is wrong: %v", teamNames)
	}

	if rleCount != 1 {
		t.Fatal("did not receive the expected number of github.RateLimitError")
	}
}
