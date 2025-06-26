package mock

import (
	"context"
	"testing"

	"github.com/google/go-github/v73/github"
)

func TestRepoGetContents(t *testing.T) {
	cases := []struct {
		name              string
		repositoryContent github.RepositoryContent
	}{
		{
			name: "fileWithoutForwardSlash",
			repositoryContent: github.RepositoryContent{
				Encoding: github.Ptr("base64"),
				Path:     github.Ptr("README.md"),
				Content:  github.Ptr("fake-content"),
			},
		},
		{
			name: "fileWithForwardSlash",
			repositoryContent: github.RepositoryContent{
				Encoding: github.Ptr("base64"),
				Path:     github.Ptr("path/test-file.txt"),
				Content:  github.Ptr("fake-content"),
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

func TestRepoGetContentsForDirectory(t *testing.T) {
	cases := []struct {
		name              string
		path              string
		repositoryContent []*github.RepositoryContent
	}{
		{
			name: "rootDirectory",
			path: "",
			repositoryContent: []*github.RepositoryContent{
				{
					Name:        github.Ptr("a.yaml"),
					DownloadURL: github.Ptr("https://raw.githubusercontent.com/foo/bar/a.yaml"),
				},
				{
					Name:        github.Ptr("b.json"),
					DownloadURL: github.Ptr("https://raw.githubusercontent.com/foo/bar/b.json"),
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(rc []*github.RepositoryContent) func(t *testing.T) {
			return func(t *testing.T) {
				mockedHTTPClient := NewMockedHTTPClient(
					WithRequestMatch(
						GetReposContentsByOwnerByRepoByPath,
						rc,
					),
				)

				client := github.NewClient(mockedHTTPClient)

				ctx := context.Background()

				_, dirContent, _, err := client.Repositories.GetContents(
					ctx,
					"foo",
					"bar",
					c.path,
					&github.RepositoryContentGetOptions{},
				)

				if len(dirContent) != len(rc) {
					t.Errorf(
						"dirContent length is %d, want %d",
						len(dirContent),
						len(rc),
					)
				}

				for i := range rc {
					if rc[i] == dirContent[1] {
						t.Errorf(
							"dirContent element at %d is '%s', want '%s'",
							i,
							*dirContent[i].Name,
							*rc[i].Name,
						)
					}
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
				Ref: github.Ptr("refs/heads/new-branch"),
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
			Ref:    github.Ptr("refs/heads/new-branch"),
			Object: &github.GitObject{SHA: github.Ptr("fake-sha")},
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
				Ref: github.Ptr("refs/heads/new-branch"),
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
