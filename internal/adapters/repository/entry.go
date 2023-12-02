package repository

import (
	"database/sql"

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
	query := `INSERT INTO entries
	(title, content, user_id, status, configs, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING *;`

	row := er.db.QueryRow(query, entry.Title, entry.Content, entry.UserID, entry.Status, entry.Configs, entry.CreatedAt, entry.UpdatedAt)

	var createdEntry domain.Entry
	err := row.Scan(&createdEntry.ID, &createdEntry.Title, &createdEntry.Content, &createdEntry.UserID, &createdEntry.Status, &createdEntry.Configs, &createdEntry.CreatedAt, &createdEntry.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &createdEntry, nil
}

func (er *EntryRepository) List(userId, skip, limit uint64, date string) ([]domain.Entry, error) {
	var entry domain.Entry
	var entries []domain.Entry

	var rows *sql.Rows
	var err error
	if len(date) > 0 {
		query := `SELECT *
		FROM entries
		WHERE user_id = $1 AND created_at::date = $2
		OFFSET $3 LIMIT $4;`

		rows, err = er.db.Query(query, userId, date, skip, limit)
		if err != nil {
			return nil, err
		}
	} else {
		query := `SELECT *
		FROM entries
		WHERE user_id = $1
		OFFSET $2 LIMIT $3;`

		rows, err = er.db.Query(query, userId, skip, limit)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&entry.ID,
			&entry.Title,
			&entry.Content,
			&entry.UserID,
			&entry.Status,
			&entry.Configs,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func (er *EntryRepository) Find(userId, id uint64) (*domain.Entry, error) {
	query := `SELECT * FROM entries WHERE id = $1 AND user_id = $2;`

	row := er.db.QueryRow(query, id, userId)

	var entry domain.Entry
	err := row.Scan(&entry.ID, &entry.Title, &entry.Content, &entry.UserID, &entry.Status, &entry.Configs, &entry.CreatedAt, &entry.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (er *EntryRepository) Update(entry *domain.Entry) (*domain.Entry, error) {
	title := nullString(entry.Title)
	content := nullString(entry.Content)
	status := nullString(entry.Status)
	configs := entry.Configs

	query := `UPDATE entries SET
	title = COALESCE($1, title),
	content = COALESCE($2, content),
	status = COALESCE($3, status),
	configs = COALESCE($4, configs),
	updated_at = $5
	WHERE id = $6 AND user_id = $7
	RETURNING *;`

	err := er.db.QueryRow(
		query,
		title,
		content,
		status,
		configs,
		entry.UpdatedAt,
		entry.ID,
		entry.UserID,
	).Scan(
		&entry.ID,
		&entry.Title,
		&entry.Content,
		&entry.UserID,
		&entry.Status,
		&entry.Configs,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (er *EntryRepository) Delete(userId, id uint64) error {
	query := `DELETE FROM entries
	WHERE id = $1 AND user_id = $2;`

	_, err := er.db.Exec(query, id, userId)
	if err != nil {
		return err
	}

	return nil
}
