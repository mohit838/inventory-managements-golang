package models

import "time"

type Test struct {
	ID          int       `db:"id" json:"id"`
	Username    string    `db:"username" json:"username"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	Role        string    `db:"role" json:"role"`
	ProfileType string    `db:"profile_type" json:"profile_type"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
