package books

import (
	"net/http"
	"start-with-crud/models"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetBooks(ctx *gin.Context) {
	var books []models.Book
	if result := h.db.Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}
