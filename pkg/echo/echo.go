package customEcho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"notes.com/pkg/templates"
)

func GetEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = templates.NewTemplate()
	return e
}
