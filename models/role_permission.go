package models

import "time"

type RolePermission struct {
	ID           int64     `db:"id" json:"id"`
	RoleID       int64     `db:"role_id" json:"role_id"`
	PermissionID int64     `db:"permission_id" json:"permission_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
