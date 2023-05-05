package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/inancgumus/effective-go/ch08/httpio"
	"github.com/inancgumus/effective-go/ch08/shortener"
)

func main() {
	var (
		addr    = flag.String("addr", "localhost:8080", "server address")
		timeout = flag.Duration("timeout", 10*time.Second, "server timeout per request")
	)
	flag.Parse()

	logger := log.New(os.Stderr, "shortener: ", log.LstdFlags|log.Lmsgprefix)
	logger.Println("starting the server on", *addr)

	shortenerServer := &shortener.Server{}
	shortenerServer.RegisterRoutes()

	server := &http.Server{
		Addr:        *addr,
		Handler:     http.TimeoutHandler(shortenerServer, *timeout, "timeout"),
		ReadTimeout: *timeout,
	}
	if os.Getenv("BITE_DEBUG") == "1" {
		server.ErrorLog = logger
		server.Handler = httpio.LoggingMiddleware(server.Handler)
	}
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		logger.Println("server closed unexpectedly:", err)
	}
}
