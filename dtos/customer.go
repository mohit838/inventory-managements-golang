package dtos

type CreateCustomerDTO struct {
	Name    string  `json:"name" binding:"required"`
	Email   *string `json:"email,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Address *string `json:"address,omitempty"`
	Notes   *string `json:"notes,omitempty"`
}

type UpdateCustomerDTO struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Address  *string `json:"address,omitempty"`
	Notes    *string `json:"notes,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}

type CustomerResponseDTO struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Email     *string `json:"email,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
	Notes     *string `json:"notes,omitempty"`
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
