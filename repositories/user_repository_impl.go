package repositories

import (
	"fmt"
	"jwt-auth-service/models/domain"
	"strings"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepositoryInterfaces {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (repository *UserRepositoryImpl) Find(email string) (*domain.User, error) {

	user := &domain.User{}
	row := repository.DB.Find(&user, "email = ?", strings.ToLower(email))
	if row.RowsAffected < 1 {
		return nil, fmt.Errorf("user tidak ditemukan")
	}

	return user, nil

}
