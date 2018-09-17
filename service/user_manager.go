package service

import "github.com/fngomez/go-twitter/domain"

var users []*domain.User

func RegisterUser(user *domain.User) {
	users = append(users, user)
	user.Id = len(users) - 1
}

func IsAuth(user *domain.User) bool {

	var userInArray = searchById(user.Id)

	return IsSameUser(user, userInArray);

	return true;
}

// Returns user pointer, if not, returns nil.
func searchById(id int) *domain.User{

	if( id < 0 || id > len(users)-1 ){
		return nil
	}

	return users[id]
}

func IsSameUser(userOne, userTwo *domain.User) bool {
	 return userOne.Email == userTwo.Email
}
