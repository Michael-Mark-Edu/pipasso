package util

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/argon2"
)

func DecryptStore(master_username, master_password string) map[string]map[string]string {
	db, err := sql.Open("sqlite3", GetDB())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
    SELECT argon_version, argon_memory, argon_time, argon_threads, decode_salt, store
    FROM accounts
    WHERE master_username == $1
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(master_username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		log.Fatal("Could not find master username in db")
	}

	var s_version, s_memory, s_time, s_threads, decode_salt, store string
	rows.Scan(&s_version, &s_memory, &s_time, &s_threads, &decode_salt, &store)

	version, err := strconv.Atoi(s_version)
	if err != nil {
		log.Fatal(err)
	}
	if version != argon2.Version {
		log.Fatal("Argon2 version desync!")
	}

	memory, err := strconv.Atoi(s_memory)
	if err != nil {
		log.Fatal(err)
	}
	time, err := strconv.Atoi(s_time)
	if err != nil {
		log.Fatal(err)
	}
	threads, err := strconv.Atoi(s_threads)
	if err != nil {
		log.Fatal(err)
	}

	key := argon2.IDKey([]byte(master_password), []byte(decode_salt), uint32(time), uint32(memory), uint8(threads), 32)

	decrypted := Decrypt([]byte(store), key)

	var data map[string]map[string]string
	err = json.Unmarshal(decrypted, &data)
	if err != nil {
		log.Fatalf("Unmarshal error: %s", err)
	}

	return data
}
