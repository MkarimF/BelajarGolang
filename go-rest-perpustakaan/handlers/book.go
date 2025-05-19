package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var books = []Book{
	{ID: "1", Title: "Golang Dasar", Author: "Budi", Year: 2020},
	{ID: "2", Title: "REST API Go", Author: "Sari", Year: 2022},
}

func GetBooks(c *gin.Context) {
	httpstatusOK(c)
}

func httpstatusOK(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "data": books})
}
func CreateBook(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}
