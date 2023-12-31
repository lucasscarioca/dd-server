package port

import "github.com/lucasscarioca/dinodiary/internal/core/domain"

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserById(id uint64) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	ListUsers(skip, limit uint64) ([]domain.PubUser, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(id uint64) error
}

type UserService interface {
	List(skip, limit uint64) ([]domain.PubUser, error)
	Find(id uint64) (*domain.SafeUser, error)
	Update(user *domain.User) (*domain.SafeUser, error)
	Delete(id uint64) error
}
