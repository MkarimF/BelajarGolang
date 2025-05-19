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

func GetBooksByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range books {
		if a.ID == id {
			httpstatusidOk(c, a)
			return
		}

	}
	httpstatusidNotFound(c, Book{ID: id})
}
func httpstatusOK(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "data": books})
}
func httpstatusidNotFound(c *gin.Context, books Book) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id u looking for doesn't exist", "data": books.ID})
}

func httpstatusidOk(c *gin.Context, book Book) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ada tuh", "data": book.ID})
}

//	func httpstatusNotFound(c *gin.Context) {
//		id := c.Param("id")
//		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id u looking for doesn't exist", "data": id})
//	}
func CreateBook(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}
