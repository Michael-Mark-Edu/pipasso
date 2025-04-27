package main

import (
	"TODO/internal/cli"
	"TODO/internal/util"
	"log"
	"os"
	"strings"
)

func main() {
	argv := os.Args
	argc := len(argv[1:])

	if argc <= 0 {
		log.Fatal("<help message>\n")
	}

	switch strings.ToLower(argv[1]) {
	case "init":
		cli.Init()
	case "add-account":
		if argc < 3 {
			log.Fatal("Not enough params")
		}
		cli.AddAccount(argv[2], argv[3])
	case "verify":
		if argc < 3 {
			log.Fatal("Not enough params")
		}
		log.Println(util.Verify(argv[2], argv[3]))
	case "add":
		if argc < 6 {
			log.Fatal("Not enough params")
		}
		cli.Add(argv[2], argv[3], argv[4], argv[5], argv[6])
	case "decrypt":
		if argc < 3 {
			log.Fatal("Not enough params")
		}
		log.Println(util.DecryptStore(argv[2], argv[3]))
	case "list":
		if argc < 3 {
			log.Fatal("Not enough params")
		}
		cli.List(argv[2], argv[3])
	default:
		log.Fatal("Invalid parameter: " + argv[1] + "\n")
	}
}
