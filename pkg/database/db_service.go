package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	err := CreateNotesTable()
	if err != nil {
		log.Fatal(err)
	}

	err = CreateUserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = AddNoteToDB("1", "Note 1", "This is the first note.", 1)
	if err != nil {
		log.Fatal(err)
	}

	err = AddNoteToDB("2", "Note 2", "This is the second note.", 1)
	if err != nil {
		log.Fatal(err)
	}

	err = AddNoteToDB("3", "Note 3", "This is the third note.", 1)
	if err != nil {
		log.Fatal(err)
	}

	err = AddNoteToDB("4", "Note 4", "This is the fourth note.", 1)
	if err != nil {
		log.Fatal(err)
	}

	err = DeleteNoteFromDB(2)
	if err != nil {
		log.Fatal(err)
	}

	err = AddUserToDB("1", "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}
}
