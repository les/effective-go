package httpio

import "net/http"

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
func Error(code int, message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, message, code)
	}
}
