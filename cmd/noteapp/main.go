package main

import (
	"/pkg/echo"

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

func main() {
	e := echo.GetEcho()

	info := newInfo("Hello World", "Welcome to my note app!")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", info)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
