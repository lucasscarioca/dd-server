package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucasscarioca/dinodiary/internal/core/utils"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Entries   []Entry   `json:"entries"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        uuid.New(),
		Name:      name,
		Avatar:    "", //TODO: randomize default avatar
		Email:     email,
		Password:  hashedPassword,
		Entries:   []Entry{},
		CreatedAt: time.Now(),
	}, nil
}
