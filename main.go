package main

import (
	"jwt-auth-service/config"
	"jwt-auth-service/controllers"
	"jwt-auth-service/repositories"
	"jwt-auth-service/services"

	"github.com/labstack/echo/v4"
)

func main() {

	db := config.NewDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	e := echo.New()
	e.POST("/auth/login", userController.Login)

	e.Logger.Fatal(e.Start(":3001"))

}
