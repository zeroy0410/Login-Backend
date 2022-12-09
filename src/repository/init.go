package repository

import (
	"Login-Backend/src/model"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var Database *gorm.DB

func Initialize() error {
	var err error
	Database, err = gorm.Open(sqlite.Open("../../User.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	Database.AutoMigrate(&model.User{})
	return nil
}