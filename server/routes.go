package server

import (
	"api/api"
	"api/controller"
	"api/database"
)

func (s *Server) RegisterRoutes() {
	db := database.NewDB()

	userRepo := database.NewUserRepo(db)
	userController := controller.NewUserController(userRepo)
	userApi := api.NewUserApi(userController)
	userGrp := s.engine.Group("user")
	{
		userGrp.GET("/hello", userApi.Hello)
		userGrp.POST("", userApi.CreateUser)
		userGrp.GET("/:userName", userApi.GetUser)
	}
}
