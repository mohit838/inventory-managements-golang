package dtos

import "time"

type CreateNotificationDTO struct {
	UserID  *int64 `json:"user_id,omitempty"`
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
	Type    string `json:"type" binding:"required"`
}

type NotificationResponseDTO struct {
	ID        int64      `json:"id"`
	UserID    *int64     `json:"user_id,omitempty"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	Type      string     `json:"type"`
	IsRead    bool       `json:"is_read"`
	ReadAt    *time.Time `json:"read_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}
