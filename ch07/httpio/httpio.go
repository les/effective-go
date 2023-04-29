package httpio

import (
	"net/http"

	"github.com/inancgumus/effective-go/ch07/bite"
)

// Handler is a http.Handler that allows chaining handlers.
type Handler func(w http.ResponseWriter, r *http.Request) http.Handler

// ServeHTTP implements the http.Handler interface.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if next := h(w, r); next != nil {
		next.ServeHTTP(w, r)
	}
}

// Error returns an httpio.Handler that writes the given error message
// to the response. If the error is internal, it logs it and hides the
// actual error message from the client.
func Error(code int, message string) Handler {
	return func(w http.ResponseWriter, r *http.Request) http.Handler {
		if code == http.StatusInternalServerError {
			Log(r.Context(), "%s: %v", r.URL.Path, message)
			message = bite.ErrInternal.Error()
		}
		return JSON(code, map[string]string{
			"error": message,
		})
	}
}

// JSON returns an httpio.Handler that writes the given value to the
// response as JSON. If the value cannot be encoded, it logs the error
// and returns an internal error to the client.
func JSON(code int, v any) Handler {
	return func(w http.ResponseWriter, r *http.Request) http.Handler {
		if err := Encode(w, code, v); err != nil {
			Log(r.Context(), "%s: encoding to JSON: %v", r.URL.Path, err)
		}
		return nil
	}
}
