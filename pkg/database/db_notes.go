package database

import (
	"time"

	models "notes.com/models"
)

func AddNoteToDB(id string, title string, body string, userId int) error {
	db := GetDB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO notes (id, title, body, userId) VALUES (?, ?, ?, ?)", id, title, body, userId)
	if err != nil {
		// check for unique constraint violation and if so continue without error
		if err.Error() == "UNIQUE constraint failed: notes.id" {
			return nil
		}

		return err
	}

	return nil
}

func GetNoteFromDB(id int) (*models.Note, error) {
	db := GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT title, body FROM notes WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var title string
	var body string
	for rows.Next() {
		err = rows.Scan(&title, &body)
		if err != nil {
			return nil, err
		}
	}

	n := models.NewNote(title, body, id, 1)

	return n, nil
}

func UpdateNoteInDB(id int, title string, body string) error {
	db := GetDB()
	defer db.Close()

	_, err := db.Exec("UPDATE notes SET title = ?, body = ? WHERE id = ?", title, body, id)
	if err != nil {
		return err
	}

	return nil
}

func GetNotesFromDB(userId int) (models.Notes, error) {
	db := GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, body, softDelete FROM notes WHERE userId = ?", userId)
	if err != nil {
		return models.Notes{}, err
	}
	defer rows.Close()

	notes := models.Notes{}
	for rows.Next() {
		var id int
		var title string
		var body string
		var softDelete bool
		err = rows.Scan(&id, &title, &body, &softDelete)
		if err != nil {
			return models.Notes{}, err
		}
		if !softDelete {
			n := models.NewNote(title, body, id, userId)
			notes.Notes = append(notes.Notes, *n)
		}
	}

	return notes, nil
}

func DeleteNoteFromDB(id int) error {
	db := GetDB()
	defer db.Close()

	date := time.Now().Format("2006-01-02")

	_, err := db.Exec("UPDATE notes SET softDelete = True, dateDeleted = ? WHERE id = ?", date, id)
	if err != nil {
		return err
	}

	return nil
}

func CreateNotesTable() error {
	db := GetDB()
	defer db.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY, title TEXT, body TEXT, userId INTEGER, softDelete BOOL DEFAULT False, dateDeleted DATE DEFAULT NULL)")
	if err != nil {
		return err
	}

	return nil
}

func GetNextNoteId() (int, error) {
	db := GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT MAX(id) FROM notes")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id + 1, nil
}
