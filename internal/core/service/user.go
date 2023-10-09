package service

import (
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
