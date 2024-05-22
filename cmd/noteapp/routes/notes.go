package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	models "notes.com/models"

	database "notes.com/pkg/database"
)

var USER_ID = 1

type pageData struct {
	notes      []models.Note
	activeNote models.Note
}

func RenderNotes(c echo.Context) error {
	n, err := database.GetNotesFromDB(USER_ID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "notes", n)
}

func RenderNote(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(400, "Invalid id")
	}
	n, err := database.GetNoteFromDB(id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "editNote", n)
}

func SaveNote(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(400, "Invalid id")
	}

	title := c.FormValue("title")
	body := c.FormValue("body")

	if id == -1 {
		newId, err := database.GetNextNoteId()
		if err != nil {
			return c.String(500, "Could not get next note id")
		}

		stringId := strconv.Itoa(newId)
		err = database.AddNoteToDB(stringId, title, body, USER_ID)
		if err != nil {
			return c.String(500, "Could not add note to database")
		}

		return c.Redirect(http.StatusSeeOther, "/notes")
	} else {
		err = database.UpdateNoteInDB(id, title, body)
		if err != nil {
			return c.String(500, "Could not update note in database")
		}
		return c.Redirect(http.StatusSeeOther, "/notes")

	}
}

func NewNote(c echo.Context) error {
	newNote := models.NewNote("", "", -1, USER_ID)
	return c.Render(http.StatusOK, "editNote", newNote)
}

func DeleteNote(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(400, "Invalid id")
	}

	err = database.DeleteNoteFromDB(id)
	if err != nil {
		return c.String(500, "Could not delete note from database")
	}

	return c.Redirect(http.StatusSeeOther, "/notes")
}
