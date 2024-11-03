package server

import (
	"bookstore-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// registerRoutes register all the routes
func registerRoutes(router *gin.Engine) error {
	router.GET("/books", getAllBooksHandler)
	router.GET("/live", checkLive)
	router.GET("/books/:id", getBookHandler)
	router.POST("/books", createBookHandler)
	router.PUT("/books/:id", updateBookHandler)
	router.DELETE("/books/:id", deleteBookHandler)
	router.POST("/register", registerUser)
	router.POST("/login", loginUser)
	router.POST("/admin/users", middleware.RoleRequired("admin"), registerUser)
	return nil
}
