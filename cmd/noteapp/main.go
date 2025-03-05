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

	e.POST("/api/login", routes.Login)

	e.GET("/notes", routes.RenderNotes)

	e.POST("/notes/edit/:id", routes.RenderNote)

	e.POST("/notes/new", routes.NewNote)

	e.POST("/notes/save/:id", routes.SaveNote)

	e.DELETE("/notes/delete/:id", routes.DeleteNote)

	e.Logger.Fatal(e.Start(":8081"))
}
