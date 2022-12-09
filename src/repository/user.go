package repository

import (
	"Login-Backend/src/model"
	"Login-Backend/src/utility"
	"errors"
	// "fmt"
	// "gorm.io/gorm"
)

var UserInitFlag = false

func GetUserByAccountString(account string) (model.User, error) {
	var user model.User
	if err := Database.Model(&model.User{}).
		Where("ID = ? or nick_name = ? or email = ?", account, account, account).
		First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func ValidateUserPassword(account string, password string) (model.User, error) {
	user, err := GetUserByAccountString(account)
	if err != nil {
		return model.User{}, errors.New("account or password is incorrect")
	}
	if !utility.CheckPasswordHash(password, user.Password) {
		return model.User{}, errors.New("account or password is incorrect")
	}
	return user, nil
}

func GetUserByID(id uint) (model.User, error) {
	var user model.User
	if err := Database.Model(&model.User{}).
		Where("ID = ?",id).
		First(&user).Error;err!=nil {
			return user,err
	}
	return user,nil
}

func CreateUser(user *model.User) error {
	if UserInitFlag {
		user.Admin = true
		UserInitFlag = false
	}
	// fmt.Println(user)
	return Database.Create(user).Error
}
func UpdateUser(user *model.User) error {
	return Database.Save(user).Error
}

func DeleteUser(userID uint) error {
	return Database.Delete(&model.User{}, userID).Error
}