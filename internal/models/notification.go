package models

import "time"

type Notification struct {
	ID        int64      `db:"id" json:"id"`
	UserID    *int64     `db:"user_id" json:"user_id,omitempty"`
	Title     string     `db:"title" json:"title"`
	Message   string     `db:"message" json:"message"`
	Type      string     `db:"type" json:"type"`
	IsRead    bool       `db:"is_read" json:"is_read"`
	ReadAt    *time.Time `db:"read_at" json:"read_at,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
}
