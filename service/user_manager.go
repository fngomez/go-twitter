package service

import (
	"errors"
	"github.com/fngomez/go-twitter/domain"
)

type UserManager struct {
	RegisteredUsers []*domain.User
	LoggedUsers map[string]*domain.User
}

func NewUserManager() UserManager {
	var res = UserManager{}
	res.initializeService()
	return res
}

func (userManager *UserManager) initializeService() {
	userManager.RegisteredUsers = make([]*domain.User, 0)
	userManager.LoggedUsers = make(map[string]*domain.User, 0)
}

func (userManager *UserManager) Register(user *domain.User) {
	userManager.RegisteredUsers = append(userManager.RegisteredUsers, user)
	user.Id = len(userManager.RegisteredUsers) - 1
}

func (userManager *UserManager) IsAuth(user *domain.User) bool {

	var userInArray = userManager.searchRegisteredUserById(user.Id)

	return isSameUser(user, userInArray);

	return true;
}

// Returns user pointer, if not, returns nil.
func (userManager *UserManager) searchRegisteredUserById(id int) *domain.User{

	if id < 0 || id > len(userManager.RegisteredUsers)-1 {
		return nil
	}

	return userManager.RegisteredUsers[id]
}

func (userManager *UserManager) searchLoggedUserById(id int) *domain.User{

	if id < 0 || id > len(userManager.LoggedUsers)-1 {
		return nil
	}

	for _, user := range userManager.LoggedUsers {
		if user.Id == id {
			return user
		}
	}

	return  nil
}

func (userManager *UserManager) searchRegisteredUserByEmail(email string) *domain.User{

	for _, user := range userManager.RegisteredUsers {
		if user.Email == email {
			return user
		}
	}

	return  nil
}

func isSameUser(userOne, userTwo *domain.User) bool {
	 return userOne.Email == userTwo.Email
}

func (userManager *UserManager) Login(email, password string) error {

	var user = userManager.searchRegisteredUserByEmail(email)

	if user == nil || user.Password != password {
		return errors.New("Wrong User or Password")
	}

	userManager.LoggedUsers[email] = user

	return nil
}

func (userManager *UserManager) Logout(email string) {

	delete(userManager.LoggedUsers, email)
}
