package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	database "notes.com/pkg/database"
)

func RenderNotes(c echo.Context) error {
	n, err := database.GetNotesFromDB(1)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "notes", n)
}
