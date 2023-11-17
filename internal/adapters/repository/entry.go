package repository

import (
	"time"

	"github.com/lucasscarioca/dinodiary/internal/core/domain"
)

type EntryRepository struct {
	db *DB
}

func NewEntryRepository(db *DB) *EntryRepository {
	return &EntryRepository{
		db,
	}
}

func (er *EntryRepository) Create(entry *domain.Entry) (*domain.Entry, error) {

}

func (er *EntryRepository) List(id, skip, limit uint64, date time.Time) ([]domain.Entry, error) {

}

func (er *EntryRepository) Find(id uint64) (*domain.Entry, error) {

}

func (er *EntryRepository) Update(entry *domain.Entry) (*domain.Entry, error) {

}

func (er *EntryRepository) Delete(id uint64) error {

}
