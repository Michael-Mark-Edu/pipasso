package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Michael-Mark-Edu/pipasso/internal/cli"
)

func Start(port uint16) {
	db_dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error: Could not find the user's home directory. %s", err)
	}

	pass_dir := db_dir + string(os.PathSeparator) + ".pipasso"
	fullname := pass_dir + string(os.PathSeparator) + "password.db"
	_, err = os.Stat(fullname)
	if err != nil {
		if os.IsNotExist(err) {
			cli.Init()
		} else {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
		log.Printf("Received message from %q", r.URL.Path)
	})

	addr := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServeTLS(addr, "./tls/server.crt", "./tls/server.key", nil))
}
