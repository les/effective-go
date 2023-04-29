package shortener

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/inancgumus/effective-go/ch07/bite"
	"github.com/inancgumus/effective-go/ch07/httpio"
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
	mux.Handle(shorteningRoute, httpio.Handler(handleShorten))
	mux.Handle(resolveRoute, httpio.Handler(handleResolve))
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
func handleShorten(w http.ResponseWriter, r *http.Request) http.Handler {
	if r.Method != http.MethodPost {
		return httpio.Error(http.StatusMethodNotAllowed, "method not allowed")
	}

	var ln short.Link

	if err := httpio.Decode(http.MaxBytesReader(w, r.Body, 4_096), &ln); err != nil {
		return httpio.Error(http.StatusBadRequest, "cannot decode JSON")
	}
	if err := short.Create(r.Context(), ln); err != nil {
		return handleError(err)
	}
	return httpio.JSON(http.StatusCreated, map[string]any{
		"key": ln.Key,
	})
}

// handleResolve handles the URL resolving requests for the short links.
//
//	Status Code       Condition
//	302               The link is successfully resolved.
//	400               The request is invalid.
//	404               The link does not exist.
//	500               There is an internal error.
func handleResolve(w http.ResponseWriter, r *http.Request) http.Handler {
	key := r.URL.Path[len(resolveRoute):]

	ln, err := short.Retrieve(r.Context(), key)
	if err != nil {
		return handleError(err)
	}

	http.Redirect(w, r, ln.URL, http.StatusFound)

	return nil // success
}

// handleHealthCheck handles the health check requests.
//
//	Status Code       Condition
//	200               The server is healthy.
func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func handleError(err error) http.Handler {
	switch {
	case err == nil: // no error
		return nil
	case errors.Is(err, bite.ErrInvalidRequest):
		return httpio.Error(http.StatusBadRequest, err.Error())
	case errors.Is(err, bite.ErrExists):
		return httpio.Error(http.StatusConflict, err.Error())
	case errors.Is(err, bite.ErrNotExist):
		return httpio.Error(http.StatusNotFound, err.Error())
	default:
		return httpio.Error(http.StatusInternalServerError, err.Error())
	}
}
