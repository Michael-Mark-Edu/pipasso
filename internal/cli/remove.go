package cli

import (
	"encoding/json"
	"log"

	"github.com/Michael-Mark-Edu/pipasso/internal/util"
)

func Remove(master_username, master_password, service, username string) {
	if !util.Verify(master_username, master_password) {
		log.Fatal("Master credentials are incorrect")
	}

	var store map[string]map[string]string = util.DecryptStore(master_username, master_password)

	if username == "" {
		if _, ok := store[service]; !ok {
			log.Fatalf("Service %s does not exist.", service)
		}
		delete(store, service)
	} else {
		if _, ok := store[service]; !ok {
			log.Fatalf("Service %s does not exist.", service)
		}
		if _, ok := store[service][username]; !ok {
			log.Fatalf("Service %s does not contain username %s.", service, username)
		}
		delete(store[service], username)
		if len(store[service]) == 0 {
			delete(store, service)
		}
	}

	marshalled, err := json.Marshal(store)
	if err != nil {
		log.Fatal(err)
	}

	util.EncrpytStore(master_username, master_password, marshalled)
}
