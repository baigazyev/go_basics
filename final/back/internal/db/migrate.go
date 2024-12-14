package db

import (
	"e-commerce/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Role{},
		&models.Order{},
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.User{},
		&models.OrderItem{},
		&models.Payment{},
		&models.Review{},
		&models.ShoppingCart{},
		&models.CartItem{},
		&models.UserAddress{},
		&models.Session{},
		&models.AuditLog{},
		&models.Cache{},
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to auto-migrate models: %v", err))
	}
}
