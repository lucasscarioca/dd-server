package repository

import "github.com/lucasscarioca/dinodiary/internal/core/domain"

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
	(name, avatar, email, password)
	VALUES ($1, $2, $3, $4)
	RETURNING id, name, avatar, email, created_at;`

	row := ur.db.QueryRow(
		query,
		user.Name,
		user.Avatar,
		user.Email,
		user.Password,
	)

	var createdUser domain.User
	err := row.Scan(&createdUser.ID, &createdUser.Name, &createdUser.Avatar, &createdUser.Email, &createdUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `SELECT * FROM users WHERE email=$1`

	row := ur.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Avatar, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
