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
	createdEntry, err := es.repo.Create(entry)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return nil, port.ErrConflictingData
		}
		return nil, err
	}

	return createdEntry, nil
}

func (es *EntryService) List(userId, skip, limit uint64, date string) ([]domain.Entry, error) {
	if limit == 0 {
		limit = 10
	}

	if len(date) > 0 {
		_, err := time.Parse(time.DateOnly, date)
		if err != nil {
			return nil, port.ErrInvalidDate
		}
	}

	entries, err := es.repo.List(userId, skip, limit, date)
	if err != nil || entries == nil {
		return nil, port.ErrDataNotFound
	}

	return entries, nil
}

func (es *EntryService) Find(userId, id uint64) (*domain.Entry, error) {
	entry, err := es.repo.Find(userId, id)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	return entry, nil
}

func (es *EntryService) Update(entry *domain.Entry) (*domain.Entry, error) {
	existingEntry, err := es.repo.Find(entry.UserID, entry.ID)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	emptyData := entry.Title == "" && entry.Content == "" && entry.Status == "" && entry.Configs == nil
	sameData := existingEntry.Title == entry.Title && existingEntry.Content == entry.Content && existingEntry.Status == entry.Status
	if emptyData || sameData {
		return nil, port.ErrNoUpdatedData
	}

	newEntry, err := es.repo.Update(entry)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return nil, port.ErrConflictingData
		}
		return nil, err
	}

	return newEntry, nil
}

func (es *EntryService) Delete(userId, id uint64) error {
	_, err := es.repo.Find(userId, id)
	if err != nil {
		return port.ErrDataNotFound
	}

	err = es.repo.Delete(userId, id)
	if err != nil {
		return err
	}

	return nil
}
