package router

import (
	"github.com/mohit838/inventory-managements-golang/internal/service"
	"github.com/mohit838/inventory-managements-golang/pkg/auth"
)

type Deps struct {
	AuthService service.AuthService
	JWTService  *auth.JWTService
}
