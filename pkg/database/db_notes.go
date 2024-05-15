package database

import (
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

	_, err := db.Exec("UPDATE notes SET softDelete = True WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func CreateNotesTable() error {
	db := GetDB()
	defer db.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY, title TEXT, body TEXT, userId INTEGER, softDelete BOOL DEFAULT False)")
	if err != nil {
		return err
	}

	return nil
}
