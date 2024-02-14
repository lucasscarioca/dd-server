package port

import "github.com/lucasscarioca/dinodiary/internal/core/domain"

type DinoRepository interface {
	Create(dino *domain.Dino) (*domain.Dino, error)
	Find(userId, id uint64) (*domain.Dino, error)
	List(userId, skip, limit uint64) ([]domain.Dino, error)
	Update(dino *domain.Dino) (*domain.Dino, error)
	Delete(userId, id uint64) error
}

type DinoService interface {
	Create(dino *domain.Dino) (*domain.ParsedDino, error)
	Find(userId, id uint64) (*domain.ParsedDino, error)
	List(userId, skip, limit uint64) ([]domain.ParsedDino, error)
	Update(dino *domain.Dino) (*domain.ParsedDino, error)
	Delete(userId, id uint64) error
}
