package main

import (
	routes "notes.com/cmd/noteapp/routes"
	database "notes.com/pkg/database"
	customEcho "notes.com/pkg/echo"
)

func main() {
	database.InitDB()

	e := customEcho.GetEcho()

	e.GET("/", routes.RenderHome)

	e.GET("/login", routes.RenderLogin)

	e.GET("/notes", routes.RenderNotes)

	e.Logger.Fatal(e.Start(":42069"))
}
