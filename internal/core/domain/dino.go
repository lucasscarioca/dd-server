package domain

import (
	"encoding/json"
	"time"
)

type Dino struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Configs   []byte    `json:"configs"`
	UserID    uint64    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ParsedDino struct {
	ID        uint64                 `json:"id"`
	Name      string                 `json:"name"`
	Avatar    string                 `json:"avatar"`
	Configs   map[string]interface{} `json:"configs"`
	UserID    uint64                 `json:"userId"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

func (d *Dino) Parse() *ParsedDino {
	var parsedConfigs map[string]interface{}
	if err := json.Unmarshal(d.Configs, &parsedConfigs); err != nil {
		parsedConfigs = map[string]interface{}{}
	}

	return &ParsedDino{
		ID:        d.ID,
		Name:      d.Name,
		Avatar:    d.Avatar,
		Configs:   parsedConfigs,
		UserID:    d.UserID,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func ParseDinos(dinos []Dino) []ParsedDino {
	var res []ParsedDino
	for _, dino := range dinos {
		res = append(res, *dino.Parse())
	}
	return res
}
