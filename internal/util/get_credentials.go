package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func GetCredentials() (master_username, master_password string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Master username: ")
	master_username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	master_username = strings.Trim(master_username, "\n")

	fmt.Print("Master password: ")
	master_bytes, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	master_password = string(master_bytes)
	return
}
