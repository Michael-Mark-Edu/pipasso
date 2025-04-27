package cli

import (
	"TODO/internal/util"
	"encoding/json"
	"log"
)

func Add(master_username, master_password, service, username, password string) {
	if !util.Verify(master_username, master_password) {
		log.Fatal("Master credentials are incorrect")
	}

	var store map[string]map[string]string = util.DecryptStore(master_username, master_password)

	if _, ok := store[service]; !ok {
		store[service] = make(map[string]string)
		store[service][username] = password
	} else if _, ok := store[service][username]; !ok {
		store[service][username] = password
	} else {
		log.Fatal(service + ": " + username + " already has a password!")
	}

	marshalled, err := json.Marshal(store)
	if err != nil {
		log.Fatal(err)
	}

	util.EncrpytStore(master_username, master_password, marshalled)
}
