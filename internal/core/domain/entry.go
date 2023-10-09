package domain

import "time"

type Entry struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"userID"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
