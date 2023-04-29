package shortener

import (
	"fmt"
	"net/http"
)

const (
	shorteningRoute  = "/shorten"
	resolveRoute     = "/r/"
	healthCheckRoute = "/health"
)

// Server is a URL shortener HTTP server. Server is an http.Handler
// that can route requests to the appropriate handler.
type Server struct {
	mux *http.ServeMux
}

// RegisterRoutes registers the handlers.
func (s *Server) RegisterRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc(shorteningRoute, handleShorten)
	mux.HandleFunc(resolveRoute, handleResolve)
	mux.HandleFunc(healthCheckRoute, handleHealthCheck)
	s.mux = mux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// handleShorten handles the URL shortening requests.
//
//	Status Code       Condition
//	201               The link is successfully shortened.
//	400               The request is invalid.
//	409               The link already exists.
//	405               The request method is not POST.
//	413               The request body is too large.
//	500               There is an internal error.
func handleShorten(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "go")
}

// handleResolve handles the URL resolving requests for the short links.
//
//	Status Code       Condition
//	302               The link is successfully resolved.
//	400               The request is invalid.
//	404               The link does not exist.
//	500               There is an internal error.
func handleResolve(w http.ResponseWriter, r *http.Request) {
	const uri = "https://go.dev"
	http.Redirect(w, r, uri, http.StatusFound)
}

// handleHealthCheck handles the health check requests.
//
//	Status Code       Condition
//	200               The server is healthy.
func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
