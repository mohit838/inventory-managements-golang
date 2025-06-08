package dtos

type CreateSaleItemDTO struct {
	ProductID    int64   `json:"product_id" binding:"required"`
	Quantity     int     `json:"quantity" binding:"required,min=1"`
	UnitPrice    float64 `json:"unit_price" binding:"required"`
	DiscountRate float64 `json:"discount_rate"`
}

type SaleItemResponseDTO struct {
	ID           int64   `json:"id"`
	SaleID       int64   `json:"sale_id"`
	ProductID    int64   `json:"product_id"`
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unit_price"`
	TotalPrice   float64 `json:"total_price"`
	DiscountRate float64 `json:"discount_rate"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
