package controller

import "api/database"

type UserController struct {
	userRepo *database.UserRepo
}

func NewUserController(userRepo *database.UserRepo) *UserController {
	return &UserController{
		userRepo: userRepo,
	}
}

func toUserModel(user User) database.UserModel {
	return database.UserModel{
		UserName: user.UserName,
		Password: user.Password,
		Address:  user.Address,
	}
}

func fromUserModel(model database.UserModel) User {
	return User{
		ID:       model.ID,
		UserName: model.UserName,
		Password: model.Password,
		Address:  model.Address,
	}
}

func (u *UserController) CreateUser(user User) error {
	model := toUserModel(user)
	err := u.userRepo.CreateUser(model)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserController) GetUser(userName string) *User {
	userModel, err := u.userRepo.GetUser(userName)
	if err != nil {
		return nil
	}
	user := fromUserModel(*userModel)
	return &user
}
