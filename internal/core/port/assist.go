package port

import "github.com/lucasscarioca/dinodiary/internal/core/domain"

type AssistRepository interface {
	Create(assist *domain.Assist) (*domain.Assist, error)
	ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error)
}

type AssistService interface {
	Create(assistantId, userId uint64) error
	ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error)
}
