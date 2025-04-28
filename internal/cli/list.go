package cli

import (
	"fmt"
	"log"

	"github.com/Michael-Mark-Edu/pipasso/internal/util"
)

func List(master_username, master_password string) {
	if !util.Verify(master_username, master_password) {
		log.Fatal("Master credentials are incorrect")
	}

	var store map[string]map[string]string = util.DecryptStore(master_username, master_password)

	fmt.Println()
	for k1, v1 := range store {
		fmt.Print(k1)
		fmt.Println(":")
		for k2, v2 := range v1 {
			fmt.Print("  ")
			fmt.Print(k2)
			fmt.Print(": ")
			fmt.Println(v2)
		}
	}
}
