package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
)

type EndpointPattern = string

type MockBackendOption func(*mux.Router)

type SequencialReponseHandler struct {
	Responses [][]byte
}

type EnforceHostRoundTripper struct {
	Host                 string
	UpstreamRoundTripper http.RoundTripper
}

func (efrt *EnforceHostRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	splitHost := strings.Split(efrt.Host, "://")
	r.URL.Scheme = splitHost[0]
	r.URL.Host = splitHost[1]

	return efrt.UpstreamRoundTripper.RoundTrip(r)
}

func (srh *SequencialReponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(srh.Responses) == 0 {
		http.Error(
			w,
			fmt.Sprintf(
				"go-github-mock: not more mocks available for %s",
				r.URL.Path,
			),
			400,
		)

		return
	}

	defer func() {
		srh.Responses = srh.Responses[1:]
	}()

	w.Write(srh.Responses[0])
}

func WithRequestMatch(
	pattern EndpointPattern,
	marshalledResponses [][]byte,
) MockBackendOption {
	return func(router *mux.Router) {
		router.Handle(pattern, &SequencialReponseHandler{
			Responses: marshalledResponses,
		})
	}
}

func WithRequestMatchHandler(
	pattern EndpointPattern,
	handler func(w http.ResponseWriter, r *http.Request),
) MockBackendOption {
	return func(router *mux.Router) {
		router.HandleFunc(pattern, handler)
	}
}

func MustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)

	if err == nil {
		return b
	}

	panic(err)
}

func NewMockedHttpClient(options ...MockBackendOption) *http.Client {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("mocked response not found for %s", r.URL.Path)))
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
