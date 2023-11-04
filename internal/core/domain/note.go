package domain

import "time"

type Note struct {
	ID        uint64    `json:"id"`
	Message   string    `json:"message"`
	EntryID   uint64    `json:"entryId"`
	UserID    uint64    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
