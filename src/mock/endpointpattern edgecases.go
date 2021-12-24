package mock

import "net/http"

// GetReposZipballByOwnerByRepo is an edge case, use it if you're not specifying the
//	`ref` when querying for the `zipball`.
//
// See: https://github.com/migueleliasweb/go-github-mock/issues/18
var GetReposZipballByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/zipball",
	Method:  "GET",
}

// GetArchiveLinkHandler mitigates the edge case for getting repo archive links without
// a `{ref}`. Use this is you're not specifying a git reference when getting the link.
//
// See: https://github.com/migueleliasweb/go-github-mock/issues/18
//
// E.g.
//
// 		mockedHTTPClient := NewMockedHTTPClient(
// 			WithRequestMatchHandler(
// 				GetReposZipballByOwnerByRepo,
// 				&GetArchiveLinkHandler{
// 					Location: "http://mockmockmock/myrepo",
// 				},
// 			),
// 		)
type GetArchiveLinkHandler struct {
	Location string
}

func (alh *GetArchiveLinkHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(
		rw,
		r,
		alh.Location,
		http.StatusFound,
	)
}
