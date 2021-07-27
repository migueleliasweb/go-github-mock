package mock

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
)

// EndpointPattern models the GitHub's API endpoints
type EndpointPattern struct {
	Pattern string // eg. "/repos/{owner}/{repo}/actions/artifacts"
	Method  string // "GET", "POST", "PATCH", etc
}

// MockBackendOption is used to configure the *mux.router
// for the mocked backend
type MockBackendOption func(*mux.Router)

// FIFOReponseHandler handler implementation that
// responds to the HTTP requests following a FIFO approach.
type FIFOReponseHandler struct {
	Responses [][]byte
}

// ServeHTTP implementation of `http.Handler`
func (srh *FIFOReponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(srh.Responses) == 0 {
		panic(fmt.Sprintf(
			"go-github-mock: not more mocks available for %s",
			r.URL.Path,
		))
	}

	defer func() {
		srh.Responses = srh.Responses[1:]
	}()

	w.Write(srh.Responses[0])
}

// EnforceHostRoundTripper rewrites all requests with the given `Host`.
type EnforceHostRoundTripper struct {
	Host                 string
	UpstreamRoundTripper http.RoundTripper
}

// RoundTrip implementation of `http.RoundTripper`
func (efrt *EnforceHostRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	splitHost := strings.Split(efrt.Host, "://")
	r.URL.Scheme = splitHost[0]
	r.URL.Host = splitHost[1]

	return efrt.UpstreamRoundTripper.RoundTrip(r)
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

// NewMockedHttpClient creates and configures an http.Client that points to
// a mocked GitHub's backend API.
//
// Example:
//
// mockedHttpClient := NewMockedHttpClient(
// 		WithRequestMatch(
// 			GetUsersByUsername,
// 			[][]byte{
// 				MustMarshal(github.User{
// 					Name: github.String("foobar"),
// 				}),
// 			},
// 		),
// 		WithRequestMatch(
// 			GetUsersOrgsByUsername,
// 			[][]byte{
// 				MustMarshal([]github.Organization{
// 					{
// 						Name: github.String("foobar123thisorgwasmocked"),
// 					},
// 				}),
// 			},
// 		),
// 		WithRequestMatchHandler(
// 			GetOrgsProjectsByOrg,
// 			func(w http.ResponseWriter, _ *http.Request) {
// 				w.Write(MustMarshal([]github.Project{
// 					{
// 						Name: github.String("mocked-proj-1"),
// 					},
// 					{
// 						Name: github.String("mocked-proj-2"),
// 					},
// 				}))
// 			},
// 		),
// )

// c := github.NewClient(mockedHttpClient)
func NewMockedHttpClient(options ...MockBackendOption) *http.Client {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteError(
			w,
			http.StatusNotFound,
			fmt.Sprintf("mock response not found for %s", r.URL.Path),
		)
	})

	for _, o := range options {
		o(router)
	}

	mockServer := httptest.NewServer(router)

	c := mockServer.Client()

	c.Transport = &EnforceHostRoundTripper{
		Host:                 mockServer.URL,
		UpstreamRoundTripper: mockServer.Client().Transport,
	}

	return c
}
