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
		return nil, port.ErrDataNotFound
	}

	return users, nil
}

func (us *UserService) Find(id uint64) (*domain.User, error) {
	user, err := us.repo.GetUserById(id)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	return user, nil
}

func (us *UserService) Update(user *domain.User) (*domain.User, error) {
	existingUser, err := us.repo.GetUserById(user.ID)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	emptyData := user.Name == "" && user.Avatar == "" && user.Email == "" && user.Configs == nil
	sameData := existingUser.Name == user.Name && existingUser.Avatar == user.Avatar && existingUser.Email == user.Email && existingUser.Configs == user.Configs
	if emptyData || sameData {
		return nil, port.ErrNoUpdatedData
	}

	newUser, err := us.repo.UpdateUser(user)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return nil, port.ErrConflictingData
		}
		return nil, err
	}

	return newUser, nil
}

func (us *UserService) Delete(id uint64) error {
	err := us.repo.DeleteUser(id)
	if err != nil {
		return port.ErrDataNotFound
	}

	return nil
}
