// Package main provides a simple example of using the gateway package.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

var (
	port = flag.Int("port", -1, "specify a port")
)

func main() {
	flag.Parse()

	http.HandleFunc("/api/hello", hello)
	listener := gateway.ListenAndServe
	portStr := "n/a"

	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	log.Fatal(listener(portStr, nil))
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello, world!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, world!"))
}
