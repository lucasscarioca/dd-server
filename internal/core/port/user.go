package port

import "github.com/lucasscarioca/dinodiary/internal/core/domain"

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}

type UserService interface {
	Register(user *domain.User) (*domain.User, error)
}
