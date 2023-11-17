package service

import (
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
	"github.com/lucasscarioca/dinodiary/internal/core/utils"
)

type AuthService struct {
	repo port.UserRepository
	tp   port.TokenProvider
}

func NewAuthService(repo port.UserRepository, tp port.TokenProvider) *AuthService {
	return &AuthService{
		repo,
		tp,
	}
}

func (as *AuthService) Login(email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(email)
	if err != nil {
		return "", port.ErrInvalidCredentials
	}

	err = utils.ComparePassword(password, user.Password)
	if err != nil {
		return "", port.ErrInvalidCredentials
	}

	t, err := as.tp.Create(user.ID, email)
	if err != nil {
		return "", err
	}
	return t, nil
}

func (as *AuthService) Register(name, email, password string) (string, error) {
	user, err := domain.NewUser(name, email, password)
	if err != nil {
		return "", err
	}

	createdUser, err := as.repo.CreateUser(user)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return "", port.ErrConflictingData
		}
		return "", err
	}

	t, err := as.tp.Create(createdUser.ID, createdUser.Email)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (as *AuthService) Forgot(email string) error {
	// Find User by email

	// Create token

	// Update User with reset_token

	// Send email
	return nil
}

func (as *AuthService) Reset(password, token string) error {
	// Find User by reset_token

	// Hash new password

	// Update User and clear reset_token

	return nil
}
