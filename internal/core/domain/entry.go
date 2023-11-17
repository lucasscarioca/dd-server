package domain

import "time"

type Entry struct {
	ID        uint64         `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	UserID    uint64         `json:"userId"`
	Status    string         `json:"status"`
	Configs   map[string]any `json:"configs"` //TODO: type configs according to app
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
