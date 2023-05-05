package httpio

import (
	"encoding/json"
	"io"
	"net/http"
)

// Decode decodes the request body as JSON and stores the result
// in the given value.
func Decode(r io.Reader, v any) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

// Encode encodes the given value as JSON and writes it to the
// response.
func Encode(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
