package echo

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func GetEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = NewTemplate()
	return e
}
