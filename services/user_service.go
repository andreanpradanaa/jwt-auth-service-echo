package services

import "jwt-auth-service/models/dto"

type UserServiceInterfaces interface {
	Login(request *dto.UserLoginRequest) (*dto.UserLoginResponse, error)
}
