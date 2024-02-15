package mock

import (
	"context"
	"testing"

	"github.com/google/go-github/v59/github"
)

func TestRepoGetContents(t *testing.T) {
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

func TestPatchGitReference(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatch(
			PatchReposGitRefsByOwnerByRepoByRef,
			github.Reference{
				Ref: github.String("refs/heads/new-branch"),
			},
		),
	)

	c := github.NewClient(mockedHTTPClient)

	ctx := context.Background()

	ref, _, _ := c.Git.UpdateRef(
		ctx,
		"owner",
		"repo-name",
		&github.Reference{
			Ref:    github.String("refs/heads/new-branch"),
			Object: &github.GitObject{SHA: github.String("fake-sha")},
		},
		false,
	)

	if *(ref.Ref) != "refs/heads/new-branch" {
		t.Errorf(
			"ref.Ref is %s, want %s",
			*ref.Ref,
			"refs/heads/new-branch",
		)
	}
}

func TestGetGitReference(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatch(
			GetReposGitRefByOwnerByRepoByRef,
			github.Reference{
				Ref: github.String("refs/heads/new-branch"),
			},
		),
	)

	c := github.NewClient(mockedHTTPClient)

	ctx := context.Background()

	ref, _, _ := c.Git.GetRef(
		ctx,
		"owner",
		"repo-name",
		"refs/heads/new-branch",
	)

	if *(ref.Ref) != "refs/heads/new-branch" {
		t.Errorf(
			"ref.Ref is %s, want %s",
			*ref.Ref,
			"refs/heads/new-branch",
		)
	}
}

func TestRepositoriesGetCommitSHA1WithForwardSlash(t *testing.T) {
	mockedHTTPClient := NewMockedHTTPClient(
		WithRequestMatch(
			GetReposCommitsByOwnerByRepoByRef,
			[]byte("01234567890"),
		),
	)

	c := github.NewClient(mockedHTTPClient)

	ctx := context.Background()

	sha, _, err := c.Repositories.GetCommitSHA1(
		ctx,
		"myself",
		"mocked-repo",
		"refs/heads/ddd",
		"",
	)

	if err != nil {
		t.Errorf(
			"err is %s, want nil",
			err.Error(),
		)
	}

	if sha != "01234567890" {
		t.Errorf(
			"sha is %s, want 01234567890",
			sha,
		)
	}
}
