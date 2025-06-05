package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Michael-Mark-Edu/pipasso/internal/server"
)

func main() {
	argv := os.Args
	argc := len(argv[1:])

	if argc > 0 {
		port, err := strconv.Atoi(argv[1])
		if err != nil {
			log.Fatal(err)
		}
		uport := uint16(port)
		server.Start(uport)
	} else {
		server.Start(443)
	}
}
