package mock

import (
	"net/http"

	"github.com/gorilla/mux"
)

// WithRequestMatchHandler implements a request callback
// for the given `pattern`.
//
// For custom implementations, this handler usage is encouraged.
//
// Example:
//
// 	WithRequestMatchHandler(
// 		GetOrgsProjectsByOrg,
// 		func(w http.ResponseWriter, _ *http.Request) {
// 			w.Write(MustMarshal([]github.Project{
// 				{
// 					Name: github.String("mocked-proj-1"),
// 				},
// 				{
// 					Name: github.String("mocked-proj-2"),
// 				},
// 			}))
// 		},
// 	)
func WithRequestMatchHandler(
	ep EndpointPattern,
	handler http.Handler,
) MockBackendOption {
	return func(router *mux.Router) {
		router.Handle(ep.Pattern, handler).Methods(ep.Method)
	}
}

// WithRequestMatch implements a simple FIFO for requests
// of the given `pattern`.
//
// Once all responses have been used, it shall panic()!
//
// Example:
//
// 	WithRequestMatch(
// 		GetUsersByUsername,
// 		[][]byte{
// 			MustMarshal(github.User{
// 				Name: github.String("foobar"),
// 			}),
// 		},
// 	)
func WithRequestMatch(
	ep EndpointPattern,
	responsesFIFO ...interface{},
) MockBackendOption {
	responses := [][]byte{}

	for _, r := range responsesFIFO {
		responses = append(responses, MustMarshal(r))
	}

	return WithRequestMatchHandler(ep, &FIFOReponseHandler{
		Responses: responses,
	})
}

// WithRequestMatchPages honors pagination directives.
//
// Each page can be called multiple times.
//
// The order in which the pages are requested doesn't matter
//
// E.g.
//
// 		mockedHTTPClient := NewMockedHTTPClient(
// 			WithRequestMatchPages(
// 				GetOrgsReposByOrg,
// 				[][]byte{
// 					MustMarshal([]github.Repository{
// 						{
// 							Name: github.String("repo-A-on-first-page"),
// 						},
// 						{
// 							Name: github.String("repo-B-on-first-page"),
// 						},
// 					}),
// 					MustMarshal([]github.Repository{
// 						{
// 							Name: github.String("repo-C-on-second-page"),
// 						},
// 						{
// 							Name: github.String("repo-D-on-second-page"),
// 						},
// 					}),
// 				},
// 			),
// 		)
func WithRequestMatchPages(
	ep EndpointPattern,
	pages ...interface{},
) MockBackendOption {
	p := [][]byte{}

	for _, r := range pages {
		p = append(p, MustMarshal(r))
	}

	return WithRequestMatchHandler(ep, &PaginatedReponseHandler{
		ResponsePages: p,
	})
}
