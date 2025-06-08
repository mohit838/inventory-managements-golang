package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohit838/inventory-managements-golang/dtos"
	"github.com/mohit838/inventory-managements-golang/internal/service"
)

type AuthHandler struct {
	authSvc service.AuthService
}

func NewAuthHandler(authSvc service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

// Register godoc
// @Summary Register new user
// @Description Creates a new user and returns access/refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dtos.RegisterRequestDTO true "New user"
// @Success 200 {object} dtos.AuthResponseDTO
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var dto dtos.RegisterRequestDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.authSvc.Register(c.Request.Context(), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Login godoc
// @Summary Login user
// @Description Authenticates a user and returns access/refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body dtos.LoginRequestDTO true "User credentials"
// @Success 200 {object} dtos.AuthResponseDTO
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var dto dtos.LoginRequestDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.authSvc.Login(c.Request.Context(), &dto)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
