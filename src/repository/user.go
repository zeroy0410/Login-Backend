package repository

import (
	"Login-Backend/src/model"
	// "Login-Backend/src/utility"
	"errors"
)

func ValidateUserPassword(account string, password string) (model.User, error) {
	flag := (account == "2020150384" && password == "123456")
	if !flag {
		return model.User{}, errors.New("account or password is incorrect")
	}
	return model.Zeroy, nil
}