package service

import (
	"errors"
	"github.com/fngomez/go-twitter/domain"
)

var registeredUsers []*domain.User

var loggedUsers []*domain.User

func RegisterUser(user *domain.User) {
	registeredUsers = append(registeredUsers, user)
	user.Id = len(registeredUsers) - 1
}

func IsAuth(user *domain.User) bool {

	var userInArray = searchRegisteredUserById(user.Id)

	return IsSameUser(user, userInArray);

	return true;
}

// Returns user pointer, if not, returns nil.
func searchRegisteredUserById(id int) *domain.User{

	if id < 0 || id > len(registeredUsers)-1 {
		return nil
	}

	return registeredUsers[id]
}

func searchLoggedUserById(id int) *domain.User{

	if id < 0 || id > len(loggedUsers)-1 {
		return nil
	}

	for _, user := range loggedUsers {
		if user.Id == id {
			return user
		}
	}

	return  nil
}

func searchLoggedUserByEmail(email string) *domain.User{

	for _, user := range loggedUsers {
		if user.Email == email {
			return user
		}
	}

	return  nil
}

func IsSameUser(userOne, userTwo *domain.User) bool {
	 return userOne.Email == userTwo.Email
}

func Login(email, password string) error {

	var ptrUserLogged = searchLoggedUserByEmail(email)

	if ptrUserLogged == nil && ptrUserLogged.Password != password{
		return errors.New("Wrong User or Password")
	}

	//TODO: terminar
	return nil

}
