package port

import "github.com/lucasscarioca/dinodiary/internal/core/domain"

type AssistRepository interface {
	Create(assist *domain.Assist) (*domain.Assist, error)
	Update(assistantId, userId uint64, status bool) (*domain.Assist, error)
	ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistantsRequests(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistedUsersRequests(id, skip, limit uint64) ([]domain.PubUser, error)
	Find(assistantId, userId uint64) (*domain.Assist, error)
	FindRequest(assistantId, userId, requestedTo uint64) (*domain.Assist, error)
	Delete(assistantId, userId uint64) error
}

type AssistService interface {
	CreateRequest(assistantId, userId, createdBy uint64) error
	AcceptRequest(assistantId, userId, requestedTo uint64) error
	ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistantsRequests(id, skip, limit uint64) ([]domain.PubUser, error)
	ListAssistedUsersRequests(id, skip, limit uint64) ([]domain.PubUser, error)
	Delete(assistantId, userId uint64) error
}
