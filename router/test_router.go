package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mohit838/inventory-managements-golang/internal/handler"
	"github.com/mohit838/inventory-managements-golang/internal/service"
)

func TestRouters(r *gin.RouterGroup, testService service.TestService) {
	testHandler := handler.NewTestHandler(testService)

	r.GET("/all", testHandler.FetchAllDataLists)
}
