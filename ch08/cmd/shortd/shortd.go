package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/inancgumus/effective-go/ch08/httpio"
	"github.com/inancgumus/effective-go/ch08/short"
	"github.com/inancgumus/effective-go/ch08/shortener"
	"github.com/inancgumus/effective-go/ch08/sqlx"
)

func main() {
	var (
		addr    = flag.String("addr", "localhost:8080", "server address")
		timeout = flag.Duration("timeout", 10*time.Second, "server timeout per request")
		dns     = flag.String("db", "file:bite.db?mode=rwc", "database connection string")
	)
	flag.Parse()

	logger := log.New(os.Stderr, "shortener: ", log.LstdFlags|log.Lmsgprefix)
	logger.Println("starting the server on", *addr)

	db, err := sqlx.Dial(context.Background(), sqlx.DefaultDriver, *dns)
	if err != nil {
		logger.Println("connecting to database:", err)
		return
	}
	svc := &shortener.Service{
		LinkStore: &short.LinkStore{
			DB: db,
		},
	}
	shortenerServer := &shortener.Server{
		Service: svc,
	}
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
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Println("server closed unexpectedly:", err)
	}
}
