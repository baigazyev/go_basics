package server

import "github.com/gin-gonic/gin"

// registerRoutes register all the routes
func registerRoutes(router *gin.Engine) error {
	router.GET("/books", getAllBooksHandler)
	router.GET("/live", checkLive)
	router.GET("/books/:id", getBookHandler)
	router.POST("/books", createBookHandler)
	router.PUT("/books/:id", updateBookHandler)
	router.DELETE("/books/:id", deleteBookHandler)
	return nil
}
