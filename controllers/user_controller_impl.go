package controllers

import (
	"fmt"
	"jwt-auth-service/config"
	"jwt-auth-service/models/dto"
	"jwt-auth-service/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	service services.UserServiceInterfaces
}

func NewUserController(service services.UserServiceInterfaces) UserControllerInterfaces {
	return &UserControllerImpl{
		service: service,
	}
}

func (c *UserControllerImpl) Login(e echo.Context) error {

	userLoginRequest := dto.UserLoginRequest{}

	err := e.Bind(&userLoginRequest)
	fmt.Println(userLoginRequest)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, dto.WebResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    "",
		})
	}

	response, err := c.service.Login(&userLoginRequest)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, dto.WebResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    "",
		})
	}

	cfg, _ := config.LoadConfig("./app.env")
	duration := time.Duration(cfg.AccessTokenMaxAge) * time.Minute

	// Create a time.Time object with the calculated duration
	timeObject := time.Now().Add(duration)

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = response.Token
	cookie.Path = "/"
	cookie.Domain = "localhost"
	cookie.Expires = timeObject
	e.SetCookie(cookie)

	return e.JSON(http.StatusAccepted, dto.WebResponse{
		Status:  "succes",
		Message: "Login Sukses",
		Data:    response,
	})

}
