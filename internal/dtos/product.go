package dtos

type CreateProductDTO struct {
	Name          string  `json:"name" binding:"required"`
	Description   *string `json:"description,omitempty"`
	Price         float64 `json:"price" binding:"required"`
	StockQuantity int     `json:"stock_quantity" binding:"required"`
	CategoryID    *int64  `json:"category_id,omitempty"`
	IsFeatured    bool    `json:"is_featured"`
	IsDiscounted  bool    `json:"is_discounted"`
	DiscountRate  float64 `json:"discount_rate"`
	AlertAtStock  int     `json:"alert_at_stock"`
}

type UpdateProductDTO struct {
	Name          *string  `json:"name,omitempty"`
	Description   *string  `json:"description,omitempty"`
	Price         *float64 `json:"price,omitempty"`
	StockQuantity *int     `json:"stock_quantity,omitempty"`
	CategoryID    *int64   `json:"category_id,omitempty"`
	IsActive      *bool    `json:"is_active,omitempty"`
	IsFeatured    *bool    `json:"is_featured,omitempty"`
	IsDiscounted  *bool    `json:"is_discounted,omitempty"`
	DiscountRate  *float64 `json:"discount_rate,omitempty"`
	AlertAtStock  *int     `json:"alert_at_stock,omitempty"`
}

type ProductResponseDTO struct {
	ID            int64   `json:"id"`
	Barcode       string  `json:"barcode"`
	Name          string  `json:"name"`
	Description   *string `json:"description,omitempty"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	IsActive      bool    `json:"is_active"`
	IsDeleted     bool    `json:"is_deleted"`
	IsFeatured    bool    `json:"is_featured"`
	IsDiscounted  bool    `json:"is_discounted"`
	DiscountRate  float64 `json:"discount_rate"`
	AlertAtStock  int     `json:"alert_at_stock"`
	CategoryID    *int64  `json:"category_id,omitempty"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}
