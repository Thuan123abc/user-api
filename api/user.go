package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/controller"
)

type UserApi struct {
	userController *controller.UserController
}

func NewUserApi(userController *controller.UserController) *UserApi {
	return &UserApi{
		userController: userController,
	}
}

func (u *UserApi) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Thuan")
}

type CreateUserInput struct {
	UserName string `form:"userName"`
	Password string `form:"password"`
	Address  string `form:"address"`
}

func (u *UserApi) CreateUser(c *gin.Context) {
	var input CreateUserInput
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user := controller.User{
		UserName: input.UserName,
		Password: input.Password,
		Address:  input.Address,
	}
	err = u.userController.CreateUser(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "OK")
}

type GetUserOutput struct {
	User controller.User `json:"user"`
}

func (u *UserApi) GetUser(c *gin.Context) {
	userName := c.Param("userName")
	user := u.userController.GetUser(userName)
	c.JSON(http.StatusOK, &GetUserOutput{
		User: *user,
	})
}
