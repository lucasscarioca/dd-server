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

func (as *AssistService) Create(assistantId, userId uint64) error {
	newAssist := domain.Assist{
		AssistantId: assistantId,
		UserId:      userId,
	}

	_, err := as.repo.Create(&newAssist)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return port.ErrConflictingData
		}
		return err
	}

	return nil
}

func (as *AssistService) ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error) {
	if limit == 0 {
		limit = 10
	}

	assistants, err := as.repo.ListAssistants(id, skip, limit)
	if err != nil || assistants == nil {
		return nil, port.ErrDataNotFound
	}

	return assistants, nil
}

func (as *AssistService) ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error) {
	if limit == 0 {
		limit = 10
	}

	assistedUsers, err := as.repo.ListAssistedUsers(id, skip, limit)
	if err != nil || assistedUsers == nil {
		return nil, port.ErrDataNotFound
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
