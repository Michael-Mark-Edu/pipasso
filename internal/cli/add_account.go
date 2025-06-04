package cli

import (
	"crypto/rand"
	"database/sql"
	"log"

	"github.com/Michael-Mark-Edu/pipasso/internal/util"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/argon2"
)

func AddAccount(master_username, master_password string) {
	db, err := sql.Open("sqlite3", util.GetDB())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
    INSERT INTO accounts (
        master_username,
        argon_version,
        argon_memory,
        argon_time,
        argon_threads,
        master_salt,
        master_hash,
        decode_salt,
        store,
        remotes
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	time := 1
	memory := 64 * 1024
	threads := 4
	master_salt := make([]byte, 32)
	decode_salt := make([]byte, 32)

	rand.Read(master_salt)
	rand.Read(decode_salt)

	hash := argon2.IDKey([]byte(master_password), master_salt[:], 1, 64*1024, 4, 32)
	key := argon2.IDKey([]byte(master_password), decode_salt[:], 1, 64*1024, 4, 32)

	store := []byte("{}")
	encoded := util.Encrypt(store, key)
	servers := encoded

	_, err = stmt.Exec(master_username, argon2.Version, memory, time, threads, master_salt, hash, decode_salt, encoded, servers)
	if err != nil {
		log.Fatal(err)
	}
}
