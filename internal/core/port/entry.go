package port

import (
	"time"

	"github.com/lucasscarioca/dinodiary/internal/core/domain"
)

type EntryRepository interface {
	Create(entry *domain.Entry) (*domain.Entry, error)
	List(id, skip, limit uint64, date time.Time) ([]domain.Entry, error)
	Find(id uint64) (*domain.Entry, error)
	Update(entry *domain.Entry) (*domain.Entry, error)
	Delete(id uint64) error
}

type EntryService interface {
	Create(entry *domain.Entry) (*domain.Entry, error)
	List(id, skip, limit uint64, date time.Time) ([]domain.Entry, error)
	Find(id uint64) (*domain.Entry, error)
	Update(entry *domain.Entry) (*domain.Entry, error)
	Delete(id uint64) error
}
