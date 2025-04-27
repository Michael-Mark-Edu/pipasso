package util

import (
	"fmt"
	"log"
	"os"
)

func GetDB() string {
	db_dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(fmt.Sprintf("error: Could not find the user's home directory. %s", err))
	}

	pass_dir := db_dir + string(os.PathSeparator) + "temp_passman"
	fullname := pass_dir + string(os.PathSeparator) + "password.db"

	return fullname
}
