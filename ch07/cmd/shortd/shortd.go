package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/inancgumus/effective-go/ch07/shortener"
)

func main() {
	var (
		addr    = flag.String("addr", "localhost:8080", "server address")
		timeout = flag.Duration("timeout", 10*time.Second, "server timeout per request")
	)
	flag.Parse()

	fmt.Fprintln(os.Stderr, "starting the server on", *addr)

	shortenerServer := &shortener.Server{}
	shortenerServer.RegisterRoutes()

	server := &http.Server{
		Addr:        *addr,
		Handler:     http.TimeoutHandler(shortenerServer, *timeout, "timeout"),
		ReadTimeout: *timeout,
	}
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintln(os.Stderr, "server closed unexpectedly:", err)
	}
}
