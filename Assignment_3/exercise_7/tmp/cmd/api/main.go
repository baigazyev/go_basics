package main

import (
	"bookstore-api/internal/config"
	"bookstore-api/internal/server"
	"fmt"
	"os"
)

func main() {
	// Load configuration
	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Could not create new config: %v", err)
		os.Exit(1)
	}
	// Create a new server instance
	server, err := server.NewServer(config)
	if err != nil {
		fmt.Printf("Could not create new server: %v", err)
		os.Exit(1)
	}

	// start the server
	if err := server.Start(); err != nil {
		fmt.Printf("Could not start the server: %v", err)
		os.Exit(1)
	}
}
