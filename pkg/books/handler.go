package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	handler := &Handler{db: db}
	routes := router.Group("/api/v1/books")
	routes.GET("/", handler.GetBooks)
	routes.GET("/:id", handler.GetBookById)
	routes.POST("/", handler.AddBook)
	routes.PUT("/:id", handler.UpdateBook)
	routes.DELETE("/:id", handler.DeleteBook)
}
