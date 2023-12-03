package domain

import "time"

type Assist struct {
	AssistantId uint64    `json:"assistantId"`
	UserId      uint64    `json:"userId"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   uint64    `json:"createdBy"`
}
