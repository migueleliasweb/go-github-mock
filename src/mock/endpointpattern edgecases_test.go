package mock

import (
	"context"
	"testing"

	"github.com/google/go-github/v41/github"
)

func TestMockForRepoGetArchiveLink(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatchHandler(
			GetReposZipballByOwnerByRepo,
			&GetArchiveLinkHandler{
				Location: "http://mockmockmock/myrepo",
			},
		),
	)

	c := github.NewClient(mockedHTTPClient)

	ctx := context.Background()

	url, _, _ := c.Repositories.GetArchiveLink(
		ctx,
		"someone",
		"myrepo",
		github.Zipball,
		&github.RepositoryContentGetOptions{},
		true,
	)

	if url.Host != "mockmockmock" {
		t.Errorf("url host is %s want mockmockmock", url.Host)
	}
}
