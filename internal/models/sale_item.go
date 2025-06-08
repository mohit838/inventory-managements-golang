package models

import "time"

type SaleItem struct {
	ID           int64     `db:"id" json:"id"`
	SaleID       int64     `db:"sale_id" json:"sale_id"`
	ProductID    int64     `db:"product_id" json:"product_id"`
	Quantity     int       `db:"quantity" json:"quantity"`
	UnitPrice    float64   `db:"unit_price" json:"unit_price"`
	TotalPrice   float64   `db:"total_price" json:"total_price"`
	DiscountRate float64   `db:"discount_rate" json:"discount_rate"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
