package domain

import (
	"time"

	"github.com/lucasscarioca/dinodiary/internal/core/utils"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		// Avatar:   "", //TODO: randomize default avatar
	}, nil
}
