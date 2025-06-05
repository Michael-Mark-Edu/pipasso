package cli

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Michael-Mark-Edu/pipasso/internal/util"
	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	fullname := util.GetDB()
	pass_dir := filepath.Dir(fullname)

	if _, notexists := os.Stat(fullname); notexists == nil {
		log.Fatal("error: File " + fullname + " already exists.")
	}

	err := os.Mkdir(pass_dir, 0700)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	_, err = os.Create(fullname)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", fullname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
    CREATE TABLE accounts (
        master_username VARCHAR(255) NOT NULL PRIMARY KEY,

        argon_version TINYINT NOT NULL,
        argon_memory INT NOT NULL,
        argon_time INT NOT NULL,
        argon_threads TINYINT NOT NULL,

        master_salt BLOB(32) NOT NULL,
        master_hash BLOB(32) NOT NULL,
        decode_salt BLOB(32) NOT NULL,

        store BLOB NOT NULL,
        remotes BLOB NOT NULL
    )
    `)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully initialized password store at " + fullname)
}
