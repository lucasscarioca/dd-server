package service

import (
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
	"github.com/lucasscarioca/dinodiary/internal/core/port"
	"github.com/lucasscarioca/dinodiary/internal/core/utils"
)

type DinoService struct {
	repo port.DinoRepository
}

func NewDinoService(repo port.DinoRepository) *DinoService {
	return &DinoService{
		repo,
	}
}

func (ds *DinoService) Create(dino *domain.Dino) (*domain.ParsedDino, error) {
	createdDino, err := ds.repo.Create(dino)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return nil, port.ErrConflictingData
		}
		return nil, err
	}

	return createdDino.Parse(), nil
}

func (ds *DinoService) Find(userId, id uint64) (*domain.ParsedDino, error) {
	dino, err := ds.repo.Find(userId, id)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	return dino.Parse(), nil
}

func (ds *DinoService) List(userId, skip, limit uint64) ([]domain.ParsedDino, error) {
	if limit == 0 {
		limit = 10
	}

	dinos, err := ds.repo.List(userId, skip, limit)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	if dinos == nil {
		return []domain.ParsedDino{}, nil
	}

	return domain.ParseDinos(dinos), nil
}

func (ds *DinoService) Update(dino *domain.Dino) (*domain.ParsedDino, error) {
	existingDino, err := ds.repo.Find(dino.UserID, dino.ID)
	if err != nil {
		return nil, port.ErrDataNotFound
	}

	emptyData := dino.Name == "" && dino.Avatar == "" && dino.Configs == nil
	sameData := existingDino.Name == dino.Name && existingDino.Avatar == dino.Avatar

	if emptyData || sameData {
		return nil, port.ErrNoUpdatedData
	}

	if utils.EmptyConfigs(dino.Configs) {
		dino.Configs = existingDino.Configs
	}

	newDino, err := ds.repo.Update(dino)
	if err != nil {
		if port.IsUniqueConstraintViolationError(err) {
			return nil, port.ErrConflictingData
		}
		return nil, err
	}

	return newDino.Parse(), nil
}

func (ds *DinoService) Delete(userId, id uint64) error {
	_, err := ds.repo.Find(userId, id)
	if err != nil {
		return port.ErrDataNotFound
	}

	err = ds.repo.Delete(userId, id)
	if err != nil {
		return err
	}

	return nil
}
