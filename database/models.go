package database

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserName string `gorm:"uniqueIndex,size:255"`
	Password string
	Address  string
}

func (UserModel) TableName() string {
	return "users"
}
