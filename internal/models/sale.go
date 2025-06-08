package models

import "time"

type Sale struct {
	ID            int64      `db:"id" json:"id"`
	InvoiceNumber string     `db:"invoice_number" json:"invoice_number"`
	CustomerID    *int64     `db:"customer_id" json:"customer_id,omitempty"`
	EmployeeID    int64      `db:"employee_id" json:"employee_id"`
	OtpCode       *string    `db:"otp_code" json:"otp_code,omitempty"`
	OtpVerifiedAt *time.Time `db:"otp_verified_at" json:"otp_verified_at,omitempty"`
	TotalAmount   float64    `db:"total_amount" json:"total_amount"`
	PaidAmount    float64    `db:"paid_amount" json:"paid_amount"`
	PaymentMethod *string    `db:"payment_method" json:"payment_method,omitempty"`
	IsDeleted     bool       `db:"is_deleted" json:"is_deleted"`
	BaseModel
}
