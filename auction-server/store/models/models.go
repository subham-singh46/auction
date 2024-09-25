package models

import "time"

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Mobile    string    `json:"mobile" db:"mobile"`
	Salt      string    `json:"salt" db:"salt"`
	PwHash    string    `json:"pw_hash" db:"pw_hash"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
