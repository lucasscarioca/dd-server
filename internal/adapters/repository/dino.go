package repository

import (
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
)

type DinoRepository struct {
	db *DB
}

func NewDinoRepository(db *DB) *DinoRepository {
	return &DinoRepository{
		db,
	}
}

func (dr *DinoRepository) Create(dino *domain.Dino) (*domain.Dino, error) {
	query := `INSERT INTO dinos
	(name, avatar, configs, user_id) VALUES ($1, $2, $3, $4)
	RETURNING *;`

	row := dr.db.QueryRow(query, dino.Name, dino.Avatar, dino.Configs, dino.UserID)

	var createdDino domain.Dino
	err := row.Scan(&createdDino.ID, &createdDino.Name, &createdDino.Avatar, &createdDino.Configs, &createdDino.UserID, &createdDino.CreatedAt, &createdDino.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &createdDino, nil
}

func (dr *DinoRepository) Find(userId, id uint64) (*domain.Dino, error) {
	query := `SELECT * FROM dinos WHERE id = $1 AND user_id = $2;`

	row := dr.db.QueryRow(query, id, userId)

	var dino domain.Dino
	err := row.Scan(&dino.ID, &dino.Name, &dino.Avatar, &dino.Configs, &dino.UserID, &dino.CreatedAt, &dino.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &dino, nil
}

func (dr *DinoRepository) List(userId, skip, limit uint64) ([]domain.Dino, error) {
	var dino domain.Dino
	var dinos []domain.Dino

	query := `SELECT *
	FROM dinos
	WHERE user_id = $1
	OFFSET $2 LIMIT $3;`

	rows, err := dr.db.Query(query, userId, skip, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&dino.ID,
			&dino.Name,
			&dino.Avatar,
			&dino.Configs,
			&dino.UserID,
			&dino.CreatedAt,
			&dino.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		dinos = append(dinos, dino)
	}

	return dinos, nil
}

func (dr *DinoRepository) Update(dino *domain.Dino) (*domain.Dino, error) {
	name := nullString(dino.Name)
	avatar := nullString(dino.Avatar)
	configs := dino.Configs
	query := `UPDATE dinos SET
	name = COALESCE($1, name),
	avatar = COALESCE($2, avatar),
	configs = COALESCE($3, configs),
	updated_at = $4
	WHERE id = $5 AND user_id = $6
	RETURNING *;`

	err := dr.db.QueryRow(
		query,
		name,
		avatar,
		configs,
		dino.UpdatedAt,
		dino.ID,
		dino.UserID,
	).Scan(
		&dino.ID,
		&dino.Name,
		&dino.Avatar,
		&dino.Configs,
		&dino.UserID,
		&dino.CreatedAt,
		&dino.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return dino, nil
}

func (dr *DinoRepository) Delete(userId, id uint64) error {
	query := `DELETE FROM dinos
	WHERE id = $1 AND user_id = $2;`

	_, err := dr.db.Exec(query, id, userId)
	if err != nil {
		return err
	}

	return nil
}
