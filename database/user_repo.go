package database

import "gorm.io/gorm"

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	err := db.AutoMigrate(&UserModel{})
	if err != nil {
		return nil
	}
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(model UserModel) error {
	err := u.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetUser(userName string) (*UserModel, error) {
	var model UserModel
	err := u.db.Model(&UserModel{}).Where("user_name=?", userName).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
