package service

import "github.com/fngomez/go-twitter/domain"

//var users []domain.User

var _user *domain.User

func RegisterUser(user *domain.User) {
	_user = user
}
