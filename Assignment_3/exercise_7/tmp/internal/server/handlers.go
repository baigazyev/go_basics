package server

import (
	"bookstore-api/internal/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// checkLive checks server is running or not
func checkLive(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "running"})

}

func getAllBooksHandler(ctx *gin.Context) {
	books := book.GetAllBooks()    // Получаем срез книг
	ctx.JSON(http.StatusOK, books) // Возвращаем срез книг в формате JSON
}

func getBookHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if b, exists := book.GetBook(id); exists {
		ctx.JSON(http.StatusOK, b)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
}

// createBookHandler создает новую книгу
func createBookHandler(ctx *gin.Context) {
	var newBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   int    `json:"year"`
	}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	book := book.AddBook(newBook.Title, newBook.Author, newBook.Year)
	ctx.JSON(http.StatusCreated, book)
}

// updateBookHandler обновляет книгу по ID
func updateBookHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var updatedBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   int    `json:"year"`
	}

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if book, exists := book.UpdateBook(id, updatedBook.Title, updatedBook.Author, updatedBook.Year); exists {
		ctx.JSON(http.StatusOK, book)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
}

// deleteBookHandler удаляет книгу по ID
func deleteBookHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if book.DeleteBook(id) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
}
