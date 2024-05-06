package main

import (
	"os"
	"start-with-crud/pkg/books"
	"start-with-crud/pkg/common/db"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	conn_string := os.Getenv("CONNECTION_STRING")

	router := gin.Default()
	go func() {
		db := db.Init(conn_string)
		books.RegisterRoutes(router, db)
	}()

	router.Run(":" + port)
}
