package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RenderLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
