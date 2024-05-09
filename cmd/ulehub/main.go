// Package main is the entry point of the application.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

var (
	port = flag.String("port", "3000", "specify a port")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	log.Fatal(gateway.ListenAndServe(":"+*port, nil))
}
