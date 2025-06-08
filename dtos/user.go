package dtos

type UserRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
	RoleID   int64  `json:"role_id" binding:"required"`
}

type UserResponseDTO struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleID   int64  `json:"role_id"`
	IsActive bool   `json:"is_active"`
}
