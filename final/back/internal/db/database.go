package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "password"
	dbName     = "ecommerce"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	return db
}
