package main

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define the User model
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Age  int    `gorm:"not null"`
}

func main() {
	// Set up PostgreSQL connection info
	dsn := "host=localhost user=myuser password=password dbname=mydb2 port=5432 sslmode=disable"
	// Replace 'myuser', 'password', and 'mydb' with your own credentials

	// Connect to the database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Automatically migrate the schema to create the users table
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migration completed successfully.")

	// Insert users into the database
	insertUsers(db)

	// Query and print users from the database
	queryUsers(db)
}

// Function to insert some users into the database
func insertUsers(db *gorm.DB) {
	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatal("Failed to insert user:", result.Error)
		}
		fmt.Printf("Inserted user: %s, age: %d\n", user.Name, user.Age)
	}
}

// Function to query and print all users from the database
func queryUsers(db *gorm.DB) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		log.Fatal("Failed to query users:", result.Error)
	}

	fmt.Println("Users in the database:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}



