package database

import (
	"database/sql"
	"log"
)

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/notes_database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
