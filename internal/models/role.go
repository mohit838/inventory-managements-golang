package models

type Role struct {
	ID   int64   `db:"id" json:"id"`
	Name string  `db:"name" json:"name"`
	Code *string `db:"code" json:"code,omitempty"`
	BaseModel
}
