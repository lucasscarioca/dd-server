package service

import (
	"time"

	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
	"github.com/lucasscarioca/dinodiary/internal/core/utils"
)

type EntryService struct {
	repo port.EntryRepository
}

func NewEntryService(repo port.EntryRepository) *EntryService {
	return &EntryService{
		repo,
	}
}

func (es *EntryService) Create(entry *domain.Entry) (*domain.ParsedEntry, error) {
	createdEntry, err := es.repo.Create(entry)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return nil, port.ErrConflictingData
		}
		return nil, err
	}

	return createdEntry.Parse(), nil
}

func (es *EntryService) List(userId, skip, limit uint64, date string) ([]domain.ParsedEntry, error) {
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
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	if entries == nil {
		return []domain.ParsedEntry{}, nil
	}

	return domain.ParseEntries(entries), nil
}

func (es *EntryService) Find(userId, id uint64) (*domain.ParsedEntry, error) {
	entry, err := es.repo.Find(userId, id)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	return entry.Parse(), nil
}

func (es *EntryService) Update(entry *domain.Entry) (*domain.ParsedEntry, error) {
	existingEntry, err := es.repo.Find(entry.UserID, entry.ID)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	emptyData := entry.Title == "" && entry.Content == "" && entry.Status == "" && entry.Configs == nil
	sameData := existingEntry.Title == entry.Title && existingEntry.Content == entry.Content && existingEntry.Status == entry.Status
	if emptyData || sameData {
		return nil, port.ErrNoUpdatedData
	}

	if utils.EmptyConfigs(entry.Configs) {
		entry.Configs = existingEntry.Configs
	}

	newEntry, err := es.repo.Update(entry)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return nil, port.ErrConflictingData
		}
		return nil, err
	}

	return newEntry.Parse(), nil
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
