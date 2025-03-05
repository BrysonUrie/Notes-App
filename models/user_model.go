package models

type User struct {
	Username string
	Id       int
}

func NewUser(id int, userName string) *User {
	return &User{
		Id:       id,
		Username: userName,
	}
}
