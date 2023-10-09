package domain

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Entries   []Entry   `json:"entries"`
	CreatedAt time.Time `json:"createdAt"`
}
