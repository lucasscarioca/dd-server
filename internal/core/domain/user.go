package domain

import (
	"time"

	"github.com/lucasscarioca/dinodiary/internal/core/utils"
)

type User struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Configs    any       `json:"configs"` //TODO: type configs according to app
	ResetToken string    `json:"resetToken"`
	CreatedAt  time.Time `json:"createdAt"`
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

type SafeUser struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	Configs    any       `json:"configs"` //TODO: type configs according to app
	ResetToken string    `json:"resetToken"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (u *User) Safe() *SafeUser {
	return &SafeUser{
		ID:         u.ID,
		Name:       u.Name,
		Avatar:     u.Avatar,
		Email:      u.Email,
		Configs:    u.Configs,
		ResetToken: u.ResetToken,
		CreatedAt:  u.CreatedAt,
	}
}

type PubUser struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *User) ToPub() *PubUser {
	return &PubUser{
		ID:        u.ID,
		Name:      u.Name,
		Avatar:    u.Avatar,
		CreatedAt: u.CreatedAt,
	}
}
