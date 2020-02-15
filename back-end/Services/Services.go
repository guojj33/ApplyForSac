package Services

import (
	"../Models"
)

// func LogIn(id, password string) error {
// 	accountType, err := Models.CheckAccountInfo(id, password)
// 	if err == nil { //登录信息无误
// 		curAccountType = accountType
// 		switch accountType {
// 		case Models.AccountType_User:
// 			*curUser, _ = Models.GetUser(id)
// 			CurLoginStatus = LoginStatus_User
// 		case Models.AccountType_Admin:
// 			*curAdmin, _ = Models.GetAdmin(id)
// 			CurLoginStatus = LoginStatus_Admin
// 		}
// 	}
// 	return err
// }

// func LogOut() error {
// 	curUser = nil
// 	curAdmin = nil
// 	curAccountType = Models.AccountType_NotExist
// 	CurLoginStatus = LoginStatus_Out
// 	return nil
// }

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

// func GetAccountId() string {
// 	switch curAccountType {
// 	case Models.AccountType_User:
// 		return curUser.UserId
// 	case Models.AccountType_Admin:
// 		return curAdmin.AdminId
// 	}
// 	return ""
// }
