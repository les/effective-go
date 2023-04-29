package shortener

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/inancgumus/effective-go/ch07/bite"
	"github.com/inancgumus/effective-go/ch07/short"
)

const (
	shorteningRoute  = "/shorten"
	resolveRoute     = "/r/"
	healthCheckRoute = "/health"
)

// Server is a URL shortener HTTP server. Server is an http.Handler
// that can route requests to the appropriate handler.
type Server struct {
	http.Handler
}

// RegisterRoutes registers the handlers.
func (s *Server) RegisterRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc(shorteningRoute, handleShorten)
	mux.HandleFunc(resolveRoute, handleResolve)
	mux.HandleFunc(healthCheckRoute, handleHealthCheck)
	s.Handler = mux
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
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ln := short.Link{
		Key: r.FormValue("key"),
		URL: r.FormValue("url"),
	}
	if err := short.Create(r.Context(), ln); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(ln.Key))
}

// handleResolve handles the URL resolving requests for the short links.
//
//	Status Code       Condition
//	302               The link is successfully resolved.
//	400               The request is invalid.
//	404               The link does not exist.
//	500               There is an internal error.
func handleResolve(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len(resolveRoute):]

	ln, err := short.Retrieve(r.Context(), key)
	if err != nil {
		handleError(w, err)
		return
	}

	http.Redirect(w, r, ln.URL, http.StatusFound)
}

// handleHealthCheck handles the health check requests.
//
//	Status Code       Condition
//	200               The server is healthy.
func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case err == nil: // no error
		return
	case errors.Is(err, bite.ErrInvalidRequest):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, bite.ErrExists):
		http.Error(w, err.Error(), http.StatusConflict)
	case errors.Is(err, bite.ErrNotExist):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
