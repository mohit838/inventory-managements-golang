package models

type Product struct {
	ID            int64   `db:"id" json:"id"`
	Barcode       string  `db:"barcode" json:"barcode"`
	Name          string  `db:"name" json:"name"`
	Description   *string `db:"description" json:"description,omitempty"`
	Price         float64 `db:"price" json:"price"`
	StockQuantity int     `db:"stock_quantity" json:"stock_quantity"`
	IsDeleted     bool    `db:"is_deleted" json:"is_deleted"`
	IsFeatured    bool    `db:"is_featured" json:"is_featured"`
	IsDiscounted  bool    `db:"is_discounted" json:"is_discounted"`
	DiscountRate  float64 `db:"discount_rate" json:"discount_rate"`
	AlertAtStock  int     `db:"alert_at_stock" json:"alert_at_stock"`
	CategoryID    *int64  `db:"category_id" json:"category_id,omitempty"`
	BaseModel
}
