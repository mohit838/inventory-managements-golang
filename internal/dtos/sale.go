package dtos

import "time"

type CreateSaleDTO struct {
	CustomerID    *int64              `json:"customer_id,omitempty"`
	EmployeeID    int64               `json:"employee_id" binding:"required"`
	PaidAmount    float64             `json:"paid_amount" binding:"required"`
	PaymentMethod *string             `json:"payment_method,omitempty"`
	Items         []CreateSaleItemDTO `json:"items" binding:"required,dive"`
	OtpCode       *string             `json:"otp_code,omitempty"`
}

type SaleResponseDTO struct {
	ID            int64                 `json:"id"`
	InvoiceNumber string                `json:"invoice_number"`
	CustomerID    *int64                `json:"customer_id,omitempty"`
	EmployeeID    int64                 `json:"employee_id"`
	TotalAmount   float64               `json:"total_amount"`
	PaidAmount    float64               `json:"paid_amount"`
	PaymentMethod *string               `json:"payment_method,omitempty"`
	IsDeleted     bool                  `json:"is_deleted"`
	OtpVerifiedAt *time.Time            `json:"otp_verified_at,omitempty"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
	Items         []SaleItemResponseDTO `json:"items"`
}
