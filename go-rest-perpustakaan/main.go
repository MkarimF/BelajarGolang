package main

import (
	"go-rest-perpustakaan/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", handlers.GetBooks)
	router.POST("/books", handlers.CreateBook)

	router.Run(":8080")
}
