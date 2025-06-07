package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohit838/inventory-managements-golang/internal/dtos"
	"github.com/mohit838/inventory-managements-golang/internal/service"
)

type TestHandler struct {
	svc service.TestService
}

func NewTestHandler(svc service.TestService) *TestHandler {
	return &TestHandler{svc: svc}
}

// GET  /api/v1/tests
func (h *TestHandler) FetchAllDataLists(c *gin.Context) {
	users, err := h.svc.FetchAllData(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// map to DTOs
	resp := make([]dtos.TestResponseDTO, len(users))
	for i, u := range users {
		resp[i] = dtos.TestResponseDTO{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		}
	}
	c.JSON(http.StatusOK, resp)
}

// // POST /api/v1/tests
// func (h *TestHandler) CreateTest(c *gin.Context) {
//     var in dtos.CreateTestDTO
//     if err := c.ShouldBindJSON(&in); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // map DTO → model
//     m := models.Test{
//         Name:  in.Name,
//         Email: in.Email,
//     }
//     id, err := h.svc.CreateTest(c.Request.Context(), &m)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // map model → response DTO
//     out := dtos.TestResponseDTO{
//         ID:    id,
//         Name:  in.Name,
//         Email: in.Email,
//     }
//     c.JSON(http.StatusCreated, out)
// }
