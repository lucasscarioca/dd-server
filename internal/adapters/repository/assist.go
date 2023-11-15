package repository

import "github.com/lucasscarioca/dinodiary/internal/core/domain"

type AssistRepository struct {
	db *DB
}

func NewAssistRepository(db *DB) *AssistRepository {
	return &AssistRepository{
		db,
	}
}

func (ar *AssistRepository) Create(assist *domain.Assist) (*domain.Assist, error) {
	query := `INSERT INTO assists
	(assistant_id, user_id) VALUES ($1, $2)
	RETURNING assistant_id, user_id, created_at;`

	row := ar.db.QueryRow(query, assist.AssistantId, assist.UserId)

	var createdAssist domain.Assist
	err := row.Scan(&createdAssist.AssistantId, &createdAssist.UserId, &createdAssist.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &createdAssist, nil
}

func (ar *AssistRepository) ListAssistants(id, skip, limit uint64) ([]domain.PubUser, error) {
	var assistant domain.PubUser
	var assistants []domain.PubUser
	query := `SELECT u.id, u.name, u.avatar, u.created_at
	FROM users u
	JOIN assists a ON u.id = a.assistant_id
	WHERE a.user_id = $1 
	OFFSET $2 LIMIT $3;`

	rows, err := ar.db.Query(query, id, skip, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&assistant.ID,
			&assistant.Name,
			&assistant.Avatar,
			&assistant.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		assistants = append(assistants, assistant)
	}

	return assistants, nil
}

func (ar *AssistRepository) ListAssistedUsers(id, skip, limit uint64) ([]domain.PubUser, error) {
	var assisted domain.PubUser
	var assistedUsers []domain.PubUser
	query := `SELECT u.id, u.name, u.avatar, u.created_at
	FROM users u
	JOIN assists a ON u.id = a.user_id
	WHERE a.assistant_id = $1 
	OFFSET $2 LIMIT $3;`

	rows, err := ar.db.Query(query, id, skip, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&assisted.ID,
			&assisted.Name,
			&assisted.Avatar,
			&assisted.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		assistedUsers = append(assistedUsers, assisted)
	}

	return assistedUsers, nil
}
