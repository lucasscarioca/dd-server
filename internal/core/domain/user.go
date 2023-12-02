package domain

import (
	"encoding/json"
	"time"

	"github.com/lucasscarioca/dinodiary/internal/core/utils"
)

type User struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Configs    []byte    `json:"configs"`
	ResetToken string    `json:"resetToken"`
	CreatedAt  time.Time `json:"createdAt"`
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	configs, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		// Avatar:   "", //TODO: randomize default avatar
		Configs: configs,
	}, nil
}

type SafeUser struct {
	ID         uint64                 `json:"id"`
	Name       string                 `json:"name"`
	Avatar     string                 `json:"avatar"`
	Email      string                 `json:"email"`
	Configs    map[string]interface{} `json:"configs"`
	ResetToken string                 `json:"resetToken"`
	CreatedAt  time.Time              `json:"createdAt"`
}

func (u *User) Safe() *SafeUser {
	var parsedConfigs map[string]interface{}
	if err := json.Unmarshal(u.Configs, &parsedConfigs); err != nil {
		parsedConfigs = map[string]interface{}{}
	}

	return &SafeUser{
		ID:         u.ID,
		Name:       u.Name,
		Avatar:     u.Avatar,
		Email:      u.Email,
		Configs:    parsedConfigs,
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
