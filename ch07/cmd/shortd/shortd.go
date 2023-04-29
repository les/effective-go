package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "server address")
	flag.Parse()

	fmt.Fprintln(os.Stderr, "starting the server on", *addr)

	shortener := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello from the shortener server!")
			// w.Write([]byte("hello from the shortener server!")) // similar to above
		},
	)
	err := http.ListenAndServe(*addr, shortener)
	if !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintln(os.Stderr, "server closed unexpectedly:", err)
	}
}
