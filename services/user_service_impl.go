package services

import (
	"fmt"
	"jwt-auth-service/helper"
	"jwt-auth-service/models/dto"
	"jwt-auth-service/repositories"
)

type UserServiceImpl struct {
	repository repositories.UserRepositoryInterfaces
}

func NewUserService(repository repositories.UserRepositoryInterfaces) UserServiceInterfaces {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (service *UserServiceImpl) Login(request *dto.UserLoginRequest) (*dto.UserLoginResponse, error) {

	// cek user
	user, err := service.repository.Find(request.Email)
	if err != nil {
		return nil, err
	}

	fmt.Println("password", user)

	// validate password
	if user.Password != request.Password {
		return nil, fmt.Errorf("password anda salah")
	}

	// generate token
	token, err := helper.GenerateToken(user.Id)
	if err != nil {
		return nil, err
	}

	// create response
	response := &dto.UserLoginResponse{
		UserData: dto.UserDataResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
		Token: token,
	}

	return response, nil

}
