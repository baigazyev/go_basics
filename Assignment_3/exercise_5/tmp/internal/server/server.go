package server

import (
	"bookstore-api/internal/config"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server is a instance of server
type Server struct {
	config *config.Config
	router *gin.Engine
}

// NewServer returns a new Server instance
func NewServer(config *config.Config) (*Server, error) {
	router := gin.New()

	// Configure CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	// Use CORS middleware
	router.Use(cors.New(corsConfig))

	server := &Server{
		config: config,
		router: router,
	}

	registerRoutes(router) // Call to register your routes

	return server, nil
}

// Start starts the server
func (s *Server) Start() error {

	if err := s.router.Run(s.config.Address + ":" + s.config.Port); err != nil {

		return fmt.Errorf("could not start the server: %v", err)
	}

	return nil
}
