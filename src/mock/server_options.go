package mock

import (
	"net/http"

	"github.com/gorilla/mux"
)

// WithRequestMatchHandler implements a request callback
// for the given `pattern`.
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
	responsesFIFO [][]byte,
) MockBackendOption {
	return func(router *mux.Router) {
		router.Handle(ep.Pattern, &FIFOReponseHandler{
			Responses: responsesFIFO,
		}).Methods(ep.Method)
	}
}

// WithRequestMatchPages
func WithRequestMatchPages(
	ep EndpointPattern,
	pages [][]byte,
) MockBackendOption {
	return func(router *mux.Router) {
		router.Handle(ep.Pattern, &PaginatedReponseHandler{
			ResponsePages: pages,
		}).Methods(ep.Method)
	}
}
