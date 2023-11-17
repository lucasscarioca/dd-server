package service

import (
	"time"

	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
)

type EntryService struct {
	repo port.EntryRepository
}

func NewEntryService(repo port.EntryRepository) *EntryService {
	return &EntryService{
		repo,
	}
}

func (es *EntryService) Create(entry *domain.Entry) (*domain.Entry, error) {
}

func (es *EntryService) List(id, skip, limit uint64, date time.Time) ([]domain.Entry, error) {

}

func (es *EntryService) Find(id uint64) (*domain.Entry, error) {

}

func (es *EntryService) Update(entry *domain.Entry) (*domain.Entry, error) {

}

func (es *EntryService) Delete(id uint64) error {

}
