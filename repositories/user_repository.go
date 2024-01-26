package repositories

import "jwt-auth-service/models/domain"

type UserRepositoryInterfaces interface {
	Find(email string) (*domain.User, error)
}
