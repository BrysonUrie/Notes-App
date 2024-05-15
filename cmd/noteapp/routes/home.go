package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Info struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func newInfo(title, message string) *Info {
	return &Info{
		Title:   title,
		Message: message,
	}
}

func RenderHome(c echo.Context) error {
	info := newInfo("Home Page", "Welcome to the home page!")
	return c.Render(http.StatusOK, "index", info)
}
