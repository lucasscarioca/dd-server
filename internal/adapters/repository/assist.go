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
	(assistant_id, user_id, created_by) VALUES ($1, $2, $3)
	RETURNING assistant_id, user_id, status, created_at, created_by;`

	row := ar.db.QueryRow(query, assist.AssistantId, assist.UserId, assist.CreatedBy)

	var createdAssist domain.Assist
	err := row.Scan(&createdAssist.AssistantId, &createdAssist.UserId, &createdAssist.Status, &createdAssist.CreatedAt, &createdAssist.CreatedBy)
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
	WHERE a.user_id = $1 AND a.status = true
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
	WHERE a.assistant_id = $1 AND a.status = true
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

func (ar *AssistRepository) ListAssistantsRequests(id, skip, limit uint64) ([]domain.PubUser, error) {
	var assistant domain.PubUser
	var assistants []domain.PubUser
	query := `SELECT u.id, u.name, u.avatar, u.created_at
	FROM users u
	JOIN assists a ON u.id = a.assistant_id
	WHERE a.user_id = $1 AND a.status = false AND a.created_by <> $1
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

func (ar *AssistRepository) ListAssistedUsersRequests(id, skip, limit uint64) ([]domain.PubUser, error) {
	var assisted domain.PubUser
	var assistedUsers []domain.PubUser
	query := `SELECT u.id, u.name, u.avatar, u.created_at
	FROM users u
	JOIN assists a ON u.id = a.user_id
	WHERE a.assistant_id = $1 AND a.status = false AND a.created_by <> $1
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

func (ar *AssistRepository) Find(assistantId, userId uint64) (*domain.Assist, error) {
	query := `SELECT assistant_id, user_id, status, created_at, created_by
	FROM assists
	WHERE assistant_id = $1 AND user_id = $2;`

	row := ar.db.QueryRow(query, assistantId, userId)

	var assist domain.Assist
	err := row.Scan(&assist.AssistantId, &assist.UserId, &assist.Status, &assist.CreatedAt, &assist.CreatedBy)
	if err != nil {
		return nil, err
	}

	return &assist, nil
}

func (ar *AssistRepository) FindRequest(assistantId, userId, requestedTo uint64) (*domain.Assist, error) {
	query := `SELECT assistant_id, user_id, status, created_at, created_by
	FROM assists
	WHERE assistant_id = $1 AND user_id = $2 AND created_by <> $3;`

	row := ar.db.QueryRow(query, assistantId, userId, requestedTo)

	var assist domain.Assist
	err := row.Scan(&assist.AssistantId, &assist.UserId, &assist.Status, &assist.CreatedAt, &assist.CreatedBy)
	if err != nil {
		return nil, err
	}

	return &assist, nil
}

func (ar *AssistRepository) Update(assistantId, userId uint64, status bool) (*domain.Assist, error) {
	query := `UPDATE assists SET
	status = $1
	WHERE assistant_id = $2 AND user_id = $3
	RETURNING assistant_id, user_id, status, created_at, created_by;`

	var assist domain.Assist
	err := ar.db.QueryRow(query, status, assistantId, userId).Scan(&assist.AssistantId, &assist.UserId, &assist.Status, &assist.CreatedAt, &assist.CreatedBy)
	if err != nil {
		return nil, err
	}

	return &assist, nil
}

func (ar *AssistRepository) Delete(assistantId, userId uint64) error {
	query := `DELETE FROM assists WHERE assistant_id = $1 AND user_id = $2;`

	_, err := ar.db.Exec(query, assistantId, userId)
	if err != nil {
		return err
	}

	return nil
}
