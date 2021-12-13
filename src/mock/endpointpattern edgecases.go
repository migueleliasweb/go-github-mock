package mock

// GetReposZipballByOwnerByRepo is an edge case, use it if you're not specifying the
//	`ref` when querying for the `zipball`. See: https://github.com/migueleliasweb/go-github-mock/issues/18
var GetReposZipballByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/zipball",
	Method:  "GET",
}
