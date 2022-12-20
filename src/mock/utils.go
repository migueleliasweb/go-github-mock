package mock

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-github/v41/github"
)

// MustMarshal helper function that wraps json.Marshal
func MustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)

	if err == nil {
		return b
	}

	panic(err)
}

// WriteError helper function to write errors to HTTP handlers
func WriteError(
	w http.ResponseWriter,
	httpStatus int,
	msg string,
	errors ...github.Error,
) {
	w.WriteHeader(httpStatus)

	w.Write(MustMarshal(gitHubErrorReponse{
		Message: msg,
		Errors:  errors,
	}))
}

// A struct with fields we need from github.ErrorResponse.
// We don't use github.ErrorResponse directly because the Response field would
// also get marshaled as null and that would override the Response that
// go-github sets before unmarshaling the response body.
type gitHubErrorReponse struct {
	Message string         `json:"message"` // error message
	Errors  []github.Error `json:"errors"`  // more detail on individual errors
}
