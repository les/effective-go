// Package shortener provides URL shortening functionality.
// At the moment, it has a server and a client.
//
// # Endpoints
//
// The service provides two endpoints:
//   - /shorten for shortening URLs.
//   - /r/{key} for resolving shortened URLs.
//   - /health for health checking the service.
//
// # Debugging
//
// To debug the service, set the BITE_DEBUG environment variable to 1.
//
// $ BITE_DEBUG=1 go run ./cmd/shortd
//
// # Curl examples
//
// Shorten a URL:
//
//	$ curl -d '{"key":"go", "url":"https://go.dev"}' localhost:8080/shorten
//
// Resolve a shortened URL:
//
//	$ curl localhost:8080/r/go
//
// Health check:
//
//	$ curl localhost:8080/health
package shortener
