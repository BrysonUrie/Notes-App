package models

type Note struct {
	Title  string
	Body   string
	Id     int
	UserId int
}

func NewNote(title string, body string, id int, userId int) *Note {
	return &Note{
		Title:  title,
		Body:   body,
		Id:     id,
		UserId: userId,
	}
}

type Notes struct {
	Notes []Note
}
