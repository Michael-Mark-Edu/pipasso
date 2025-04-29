package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Michael-Mark-Edu/pipasso/internal/cli"
	"github.com/Michael-Mark-Edu/pipasso/internal/util"
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
		master_username, master_password := util.GetCredentials()
		cli.AddAccount(master_username, master_password)
		fmt.Println("Account created successfully!")
	case "add":
		if argc < 4 {
			log.Fatal("Not enough params")
		}
		master_username, master_password := util.GetCredentials()
		cli.Add(master_username, master_password, argv[2], argv[3], argv[4])
		fmt.Println("Added account " + argv[3] + " to service " + argv[2] + " successfully!")
	case "list":
		master_username, master_password := util.GetCredentials()
		cli.List(master_username, master_password, argv[2:]...)
	default:
		log.Fatal("Invalid parameter: " + argv[1] + "\n")
	}
}
