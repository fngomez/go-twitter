package domain

type User struct {
	Id int
	Name string
	Nickname string
	Email string
	Password string
}

func NewUser(_name string, _nickname string, _email string, _password string) *User {
	return &User {-1,_name, _nickname, _email, _password }
}
