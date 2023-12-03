package service

import (
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type AssistService struct {
	repo port.AssistRepository
}

func NewAssistService(repo port.AssistRepository) *AssistService {
	return &AssistService{
		repo,
	}
}

func (as *AssistService) CreateRequest(assistantId, userId, createdBy uint64) error {
	if assistantId == userId {
		return port.ErrConflictingData
	}

	newAssist := domain.Assist{
		AssistantId: assistantId,
		UserId:      userId,
		CreatedBy:   createdBy,
	}

	_, err := as.repo.Create(&newAssist)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return port.ErrConflictingData
		}
		if port.IsForeignKeyConstraintViolationError(err) {
			return port.ErrDataNotFound
		}
		return err
	}

	return nil
}

func (as *AssistService) AcceptRequest(assistantId, userId, requestedTo uint64) error {
	if assistantId == userId {
		return port.ErrConflictingData
	}

	_, err := as.repo.FindRequest(assistantId, userId, requestedTo)
	if err != nil {
		return port.ErrDataNotFound
	}

	_, err = as.repo.Update(assistantId, userId, true)
	if err != nil {
		return err
	}

	return nil
}

func (as *AssistService) ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error) {
	if limit == 0 {
		limit = 10
	}

	assistants, err := as.repo.ListAssistants(id, skip, limit)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	if assistants == nil {
		return []domain.PubUser{}, nil
	}

	return assistants, nil
}

func (as *AssistService) ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error) {
	if limit == 0 {
		limit = 10
	}

	assistedUsers, err := as.repo.ListAssistedUsers(id, skip, limit)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	if assistedUsers == nil {
		return []domain.PubUser{}, nil
	}

	return assistedUsers, nil
}

func (as *AssistService) ListAssistantsRequests(id, skip, limit uint64) ([]domain.PubUser, error) {
	if limit == 0 {
		limit = 10
	}

	assistants, err := as.repo.ListAssistantsRequests(id, skip, limit)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	if assistants == nil {
		return []domain.PubUser{}, nil
	}

	return assistants, nil
}

func (as *AssistService) ListAssistedUsersRequests(id, skip, limit uint64) ([]domain.PubUser, error) {
	if limit == 0 {
		limit = 10
	}

	assistedUsers, err := as.repo.ListAssistedUsersRequests(id, skip, limit)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	if assistedUsers == nil {
		return []domain.PubUser{}, nil
	}

	return assistedUsers, nil
}

func (as *AssistService) Delete(assistantId, userId uint64) error {
	_, err := as.repo.Find(assistantId, userId)
	if err != nil {
		return port.ErrDataNotFound
	}

	err = as.repo.Delete(assistantId, userId)
	if err != nil {
		return err
	}

	return nil
}
