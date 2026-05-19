package mock

import (
	"net/http"
	"testing"

	"github.com/google/go-github/v87/github"
)

// MustNewGHClient creates a new GitHub client for use in tests, failing the
// test immediately if client construction returns an error.
//
// If httpClient is nil, a client is created with no options.
func MustNewGHClient(t *testing.T, httpClient *http.Client) *github.Client {
	t.Helper()

	var (
		client *github.Client
		err    error
	)

	if httpClient == nil {
		client, err = github.NewClient()
	} else {
		client, err = github.NewClient(github.WithHTTPClient(httpClient))
	}

	if err != nil {
		t.Fatalf("failed to construct github.Client: %s", err)
	}

	return client
}
