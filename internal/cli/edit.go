package cli

import (
	"encoding/json"
	"log"

	"github.com/Michael-Mark-Edu/pipasso/internal/util"
)

func Edit(master_username, master_password, service, username, password string) {
	if !util.Verify(master_username, master_password) {
		log.Fatal("Master credentials are incorrect")
	}

	var store map[string]map[string]string = util.DecryptStore(master_username, master_password)

	if _, ok := store[service]; !ok {
		log.Fatalf("Service %s does not exist.", service)
	}
	if _, ok := store[service][username]; !ok {
		log.Fatalf("Service %s does not contain username %s.", service, username)
	}

	store[service][username] = password

	marshalled, err := json.Marshal(store)
	if err != nil {
		log.Fatal(err)
	}

	util.EncrpytStore(master_username, master_password, marshalled)
}
