package repository

import (
	"github.com/lucasscarioca/dinodiary/internal/core/domain"
)

type UserRepository struct {
	db *DB
}

func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users
	(name, avatar, email, password, reset_token)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, name, avatar, email, created_at;`

	row := ur.db.QueryRow(
		query,
		user.Name,
		user.Avatar,
		user.Email,
		user.Password,
		user.ResetToken,
	)

	var createdUser domain.User
	err := row.Scan(&createdUser.ID, &createdUser.Name, &createdUser.Avatar, &createdUser.Email, &createdUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func (ur *UserRepository) GetUserById(id uint64) (*domain.User, error) {
	query := `SELECT users.id, users.name, users.avatar, users.email, users.configs, users.created_at
	FROM users
	WHERE id=$1;`

	row := ur.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Avatar, &user.Email, &user.Configs, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `SELECT users.id, users.name, users.avatar, users.email, users.password, users.configs, users.reset_token, users.created_at
	FROM users
	WHERE email=$1;`

	row := ur.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Avatar, &user.Email, &user.Password, &user.Configs, &user.ResetToken, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) ListUsers(skip, limit uint64) ([]domain.PubUser, error) {
	var user domain.PubUser
	var users []domain.PubUser
	query := `SELECT users.id, users.name, users.avatar, users.created_at 
	FROM users 
	OFFSET $1 LIMIT $2;`

	rows, err := ur.db.Query(query, skip, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Avatar,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	name := nullString(user.Name)
	avatar := nullString(user.Avatar)
	email := nullString(user.Email)
	password := nullString(user.Password)
	resetToken := nullString(user.ResetToken)
	configs := user.Configs

	query := `UPDATE users SET
	name = COALESCE($1, name),
	avatar = COALESCE($2, avatar),
	email = COALESCE($3, email),
	password = COALESCE($4, password),
	reset_token = COALESCE($5, reset_token),
	configs = COALESCE($6, configs)
	WHERE id = $7
	RETURNING *;`

	err := ur.db.QueryRow(
		query,
		name,
		avatar,
		email,
		password,
		resetToken,
		configs,
		user.ID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Avatar,
		&user.Email,
		&user.Password,
		&user.Configs,
		&user.ResetToken,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) DeleteUser(id uint64) error {
	assistQuery := `DELETE FROM assists WHERE user_id = $1 OR assistant_id = $1;`
	entriesQuery := `DELETE FROM entries WHERE user_id = $1;`
	query := `DELETE FROM users WHERE id = $1;`

	_, err := ur.db.Exec(assistQuery, id)
	if err != nil {
		return err
	}

	_, err = ur.db.Query(entriesQuery, id)
	if err != nil {
		return err
	}

	_, err = ur.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
