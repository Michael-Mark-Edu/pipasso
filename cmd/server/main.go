package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
        log.Printf("Received message from %q", r.URL.Path);
	})
	log.Fatal(http.ListenAndServeTLS(":443", "./tls/server.crt", "./tls/server.key", nil))
}
