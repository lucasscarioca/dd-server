package domain

import "time"

type Assist struct {
	AssistantId uint64    `json:"assistantId"`
	UserId      uint64    `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
}
