package util

import (
	"crypto/subtle"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/argon2"
)

func Verify(username, password string) bool {
	db, err := sql.Open("sqlite3", GetDB())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
    SELECT argon_version, argon_memory, argon_time, argon_threads, master_salt, master_hash
    FROM accounts
    WHERE master_username == $1
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		log.Fatal("Could not find master username in db")
	}

	var s_version, s_memory, s_time, s_threads, salt, master_hash string
	if err = rows.Scan(&s_version, &s_memory, &s_time, &s_threads, &salt, &master_hash); err != nil {
		log.Fatal(err)
	}

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

	rehash := argon2.IDKey([]byte(password), []byte(salt[:]), uint32(time), uint32(memory), uint8(threads), 32)

	result := subtle.ConstantTimeCompare([]byte(master_hash), []byte(rehash))

	return result == 1
}
