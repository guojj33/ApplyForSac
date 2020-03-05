package Services

import (
	"../Models"
)

func LogIn(id, password string) (Models.AccountType, error) {
	accountType, err := Models.CheckAccountInfo(id, password)
	return accountType, err
}

func RegisterNewUser(userId, password, email string) (Models.User, error) {
	newUser := Models.User{
		UserId:   userId,
		Password: password,
		Email:    email,
	}
	_, err := Models.AddUser(&newUser)
	if err != nil {
		newUser = Models.User{}
	}
	return newUser, err
}

func AddCommentAsGuest(name string, content string) (int, error) {
	return Models.AddComment(name, content)
}
