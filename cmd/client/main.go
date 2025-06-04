package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Michael-Mark-Edu/pipasso/internal/cli"
	"github.com/Michael-Mark-Edu/pipasso/internal/util"
)

func main() {
	// I'm currently using self-signed TLS certs to make the web component of
	// this project work. However, without the below line, this causes a runtime
	// error. In an actual production environment, you'd want to delete this
	// line
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	http.Get("https://localhost")

	argv := os.Args
	argc := len(argv[1:])

	if argc <= 0 || (argc > 0 && argv[1] == "help") {
		fmt.Println("usage: pipasso <command>")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("  help: Prints this message.")
		fmt.Println("  init: Initializes Pipasso for your user. Must be called before anything else.")
		fmt.Println("  add-account: Creates a new master account with a master username and password.")
		fmt.Println("  add <service> <username> <password>: Adds a username-password to a specified service.")
		fmt.Println("  remove <service> [username]: Removes either a service or a username-password store.")
		fmt.Println("  edit <service> <username> <password>: Edits a username-password to a specified service.")
		fmt.Println("  list [filters...]: Lists all username-password stores. If provided, search only in the specified services.")
		if argc <= 0 {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
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
	case "remove":
		if argc < 2 {
			log.Fatal("Not enough params")
		}
		master_username, master_password := util.GetCredentials()
		var username string
		if argc == 2 {
			username = ""
		} else {
			username = argv[3]
		}
		cli.Remove(master_username, master_password, argv[2], username)
	case "edit":
		if argc < 4 {
			log.Fatal("Not enough params")
		}
		master_username, master_password := util.GetCredentials()
		cli.Edit(master_username, master_password, argv[2], argv[3], argv[4])
	case "list":
		master_username, master_password := util.GetCredentials()
		cli.List(master_username, master_password, argv[2:]...)
	default:
		log.Fatal("Invalid parameter: " + argv[1] + "\n")
	}
}
