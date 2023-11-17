package domain

import "time"

type Dino struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	Avatar    string         `json:"avatar"`
	Configs   map[string]any `json:"configs"` //TODO: type configs according to app
	UserID    uint64         `json:"userId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
