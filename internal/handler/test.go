package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohit838/inventory-managements-golang/internal/service"
)

type TestHandler struct {
	svc service.TestService
}

func NewTestHandler(svc service.TestService) *TestHandler {
	return &TestHandler{svc: svc}
}

// GET /api/v1/all
func (h *TestHandler) FetchAllDataLists(c *gin.Context) {
	users, err := h.svc.FetchAllData(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
