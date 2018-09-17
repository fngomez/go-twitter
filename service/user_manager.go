package service

import "github.com/fngomez/go-twitter/domain"

var users []*domain.User

func RegisterUser(user *domain.User) {
	users = append(users, user)
}
