package models

import "time"

type User struct {
	ID           int64      `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Email        string     `db:"email" json:"email"`
	Password     string     `db:"password" json:"-"`
	RoleID       int64      `db:"role_id" json:"role_id"`
	OtpCode      *string    `db:"otp_code" json:"otp_code,omitempty"`
	OtpExpiresAt *time.Time `db:"otp_expires_at" json:"otp_expires_at,omitempty"`
	TwoFASecret  *string    `db:"two_fa_secret" json:"-"`
	TwoFAEnabled bool       `db:"two_fa_enabled" json:"two_fa_enabled"`
	BaseModel
}
