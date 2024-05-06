package books

import (
	"net/http"
	"start-with-crud/models"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetBookById(ctx *gin.Context) {
	var book models.Book
	if result := h.db.First(&book, ctx.Param("id")); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
