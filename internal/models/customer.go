package models

type Customer struct {
	ID      int64   `db:"id" json:"id"`
	Name    string  `db:"name" json:"name"`
	Email   *string `db:"email" json:"email,omitempty"`
	Phone   *string `db:"phone" json:"phone,omitempty"`
	Address *string `db:"address" json:"address,omitempty"`
	Notes   *string `db:"notes" json:"notes,omitempty"`
	BaseModel
}
