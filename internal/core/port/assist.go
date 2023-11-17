package port

import "github.com/lucasscarioca/dinodiary/internal/core/domain"

type AssistRepository interface {
	Create(assist *domain.Assist) (*domain.Assist, error)
	ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error)
	Find(assistantId, userId uint64) (*domain.Assist, error)
	Delete(assistantId, userId uint64) error
}

type AssistService interface {
	Create(assistantId, userId uint64) error
	ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error)
	Delete(assistantId, userId uint64) error
}
