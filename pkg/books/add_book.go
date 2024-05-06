package books

import (
	"net/http"
	"start-with-crud/models"

	"github.com/gin-gonic/gin"
)

type AddBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (h Handler) AddBook(ctx *gin.Context) {
	var req AddBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	if result := h.db.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
