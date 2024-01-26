package controllers

import "github.com/labstack/echo/v4"

type UserControllerInterfaces interface {
	Login(e echo.Context) error
}
