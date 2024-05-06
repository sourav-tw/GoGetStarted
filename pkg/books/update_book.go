package books

import (
	"net/http"
	"start-with-crud/models"

	"github.com/gin-gonic/gin"
)

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (h Handler) UpdateBook(ctx *gin.Context) {
	var req UpdateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book
	if result := h.db.First(&book, ctx.Param("id")); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = req.Title
	book.Author = req.Author

	if result := h.db.Save(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
