package util

import (
	"log"
	"os"
)

func GetDB() string {
	db_dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	pass_dir := db_dir + string(os.PathSeparator) + ".pipasso"
	fullname := pass_dir + string(os.PathSeparator) + "password.db"

	return fullname
}
