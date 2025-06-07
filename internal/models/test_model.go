package models

import "time"

type Test struct {
	ID          int64       `db:"id" json:"id"`
	Name        string    `db:"username" json:"username"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"-"` // Do not expose in JSON
	IsActive    bool      `db:"is_active" json:"is_active"`
	Role        string    `db:"role" json:"role"`
	ProfileType string    `db:"profile_type" json:"profile_type"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
