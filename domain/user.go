package domain

type User struct {
	name string
	nickname string
	email string
	password string
}

func NewUser(_name string, _nickname string, _email string, _password string) *User {
	return &User {_name, _nickname, _email, _password }
}
