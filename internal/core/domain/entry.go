package domain

import (
	"encoding/json"
	"time"
)

type Entry struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint64    `json:"userId"`
	Status    string    `json:"status"`
	Configs   []byte    `json:"configs"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ParsedEntry struct {
	ID        uint64                 `json:"id"`
	Title     string                 `json:"title"`
	Content   string                 `json:"content"`
	UserID    uint64                 `json:"userId"`
	Status    string                 `json:"status"`
	Configs   map[string]interface{} `json:"configs"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

func (e *Entry) Parse() *ParsedEntry {
	var parsedConfigs map[string]interface{}
	if err := json.Unmarshal(e.Configs, &parsedConfigs); err != nil {
		parsedConfigs = map[string]interface{}{}
	}

	return &ParsedEntry{
		ID:        e.ID,
		Title:     e.Title,
		Content:   e.Content,
		UserID:    e.UserID,
		Status:    e.Status,
		Configs:   parsedConfigs,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ParseEntries(entries []Entry) []ParsedEntry {
	var res []ParsedEntry
	for _, entry := range entries {
		res = append(res, *entry.Parse())
	}
	return res
}
