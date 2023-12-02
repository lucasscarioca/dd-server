package port

import (
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
)

type EntryRepository interface {
	Create(entry *domain.Entry) (*domain.Entry, error)
	List(userId, skip, limit uint64, date string) ([]domain.Entry, error)
	Find(userId, id uint64) (*domain.Entry, error)
	Update(entry *domain.Entry) (*domain.Entry, error)
	Delete(userId, id uint64) error
}

type EntryService interface {
	Create(entry *domain.Entry) (*domain.ParsedEntry, error)
	List(userId, skip, limit uint64, date string) ([]domain.ParsedEntry, error)
	Find(userId, id uint64) (*domain.ParsedEntry, error)
	Update(entry *domain.Entry) (*domain.ParsedEntry, error)
	Delete(userId, id uint64) error
}
