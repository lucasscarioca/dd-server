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

func (us *UserService) List(skip, limit uint64) ([]domain.PubUser, error) {
	if limit == 0 {
		limit = 10
	}

	users, err := us.repo.ListUsers(skip, limit)
	if err != nil || users == nil {
		return nil, port.ErrDataNotFound
	}

	return users, nil
}

func (us *UserService) Find(id uint64) (*domain.SafeUser, error) {
	user, err := us.repo.GetUserById(id)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	return user.Safe(), nil
}

func (us *UserService) Update(user *domain.User) (*domain.SafeUser, error) {
	existingUser, err := us.repo.GetUserById(user.ID)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	emptyData := user.Name == "" && user.Avatar == "" && user.Email == "" && user.Configs == nil
	sameData := existingUser.Name == user.Name && existingUser.Avatar == user.Avatar && existingUser.Email == user.Email
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

	return newUser.Safe(), nil
}

func (us *UserService) Delete(id uint64) error {
	_, err := us.repo.GetUserById(id)
	if err != nil {
		return port.ErrDataNotFound
	}

	err = us.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
