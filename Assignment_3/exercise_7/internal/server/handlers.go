package server

import (
	"bookstore-api/internal/book"
	"bookstore-api/internal/user"
	"net/http"
	"strconv"

	"bookstore-api/internal/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// checkLive checks server is running or not
func checkLive(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "running"})

}

// ALLL HANDLERS ABOUT BOOKS
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

// ALLL HANDLERS ABOUT BOOKS

func registerUser(ctx *gin.Context) {
	var newUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // Include Role in request
	}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user := user.AddUser(newUser.Username, newUser.Password, newUser.Role)
	ctx.JSON(http.StatusCreated, user)
}

// User login handler
func loginUser(ctx *gin.Context) {
	var loginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate user credentials
	for _, u := range user.GetAllUsers() {
		if u.Username == loginUser.Username {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginUser.Password))
			if err == nil {
				token, _ := auth.GenerateJWT(u.Username)
				ctx.JSON(http.StatusOK, gin.H{"token": token})
				return
			}
			break
		}
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			ctx.Abort()
			return
		}

		username, err := auth.ValidateJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("username", username)
		ctx.Next()
	}
}
