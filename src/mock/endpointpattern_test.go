package mock

import (
	"context"
	"testing"

	"github.com/google/go-github/v41/github"
)

func TestRepoGetContents2(t *testing.T) {
	cases := []struct {
		name              string
		repositoryContent github.RepositoryContent
	}{
		{
			name: "fileWithoutForwardSlash",
			repositoryContent: github.RepositoryContent{
				Encoding: github.String("base64"),
				Path:     github.String("README.md"),
				Content:  github.String("fake-content"),
			},
		},
		{
			name: "fileWithForwardSlash",
			repositoryContent: github.RepositoryContent{
				Encoding: github.String("base64"),
				Path:     github.String("path/test-file.txt"),
				Content:  github.String("fake-content"),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(rc github.RepositoryContent) func(t *testing.T) {
			return func(t *testing.T) {
				mockedHTTPClient := NewMockedHTTPClient(
					WithRequestMatch(
						GetReposContentsByOwnerByRepoByPath,
						rc,
					),
				)

				c := github.NewClient(mockedHTTPClient)

				ctx := context.Background()

				fileContent, _, _, err := c.Repositories.GetContents(
					ctx,
					"foo",
					"bar",
					*rc.Path,
					&github.RepositoryContentGetOptions{},
				)

				if *(fileContent.Content) != *rc.Content {
					t.Errorf(
						"fileContent.Content is %s, want %s",
						*(fileContent.Content),
						*rc.Content,
					)
				}

				if err != nil {
					t.Errorf(
						"err is %s, want nil",
						err.Error(),
					)
				}
			}
		}(c.repositoryContent))
	}
}
