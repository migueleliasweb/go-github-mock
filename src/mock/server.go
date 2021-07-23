package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

type EndpointPattern = string

type MockBackendOption func(*mux.Router)

type SequencialReponseHandler struct {
	Responses [][]byte
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

func MustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)

	if err == nil {
		return b
	}

	panic(err)
}

func NewMockedHttpClient(options ...MockBackendOption) *http.Client {
	router := mux.NewRouter()

	for _, o := range options {
		o(router)
	}

	mockServer := httptest.NewServer(router)

	defer mockServer.Start()

	return mockServer.Client()
}
