package service

import (
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

func (us *UserService) List(skip, limit uint64) ([]domain.User, error) {
	users, err := us.repo.ListUsers(skip, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}
