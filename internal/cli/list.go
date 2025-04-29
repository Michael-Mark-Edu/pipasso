package cli

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/Michael-Mark-Edu/pipasso/internal/util"
)

func List(master_username, master_password string, services ...string) {
	if !util.Verify(master_username, master_password) {
		log.Fatal("Master credentials are incorrect")
	}

	var store map[string]map[string]string = util.DecryptStore(master_username, master_password)
	flag := false

	fmt.Println()
	for k1, v1 := range store {
		if len(services) > 0 && !slices.Contains(services, k1) {
			continue
		}
		flag = true
		fmt.Print(k1)
		fmt.Println(":")
		for k2, v2 := range v1 {
			fmt.Print("  ")
			fmt.Print(k2)
			fmt.Print(": ")
			fmt.Println(v2)
		}
	}

	if !flag {
		fmt.Println("Could not find any passwords associated with the provided services.")
		os.Exit(1)
	}
}
