package httpio

import (
	"context"
	"net/http"
	"time"
)

// LoggingMiddleware logs the request and response times
// using http.Server.ErrorLog if available.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Since(start)
		Log(r.Context(), "%s %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, end)
	})
}

// Log logs the given format and args using Server.ErrorLog
// if available.
func Log(ctx context.Context, format string, args ...any) {
	s, _ := ctx.Value(http.ServerContextKey).(*http.Server)
	if s == nil || s.ErrorLog == nil {
		return
	}
	s.ErrorLog.Printf(format, args...)
}
