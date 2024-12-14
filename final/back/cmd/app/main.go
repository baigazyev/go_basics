package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"e-commerce/internal/cache"
	"e-commerce/internal/db"
	"e-commerce/internal/repositories"
	"e-commerce/internal/routes"

	"github.com/redis/go-redis/v9"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	database := db.InitDatabase()
	db.MigrateDatabase(database)

	// Initialize the router
	router := mux.NewRouter()
	routes.RegisterRoutes(router, database)

	// Enable CORS
	corsRouter := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                                       // Allow all origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // Allowed headers
	)(router)

	// Start the server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Инициализация гибридного кэша
	cacheRepo := repositories.NewCacheRepository(database)
	hybridCache := cache.NewHybridCache(redisClient, cacheRepo)

	// Пример использования
	err := hybridCache.SetCache("example_key", "example_value", 10*time.Minute)
	if err != nil {
		log.Fatalf("Ошибка при установке кэша: %v", err)
	}

	value, err := hybridCache.GetCache("example_key")
	if err != nil {
		log.Fatalf("Ошибка при получении кэша: %v", err)
	}

	log.Printf("Значение из кэша: %s", value)
}
