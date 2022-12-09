package model

import (
	"gorm.io/gorm"
)

type User struct {
	// ID				uint		`gorm:"unique" json:"id"`
	Uid      string `gorm:"unique" json:"uid"`
	Admin    bool   `json:"admin"`
	Password string `json:"password"`
	NickName string `gorm:"unique" json:"nick-name"`
	Email    string `gorm:"unique" json:"email"`
	gorm.Model
}
