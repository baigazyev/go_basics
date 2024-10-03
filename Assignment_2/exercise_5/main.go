package main

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define the User and Profile models
type User struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"not null"`
	Age     int     `gorm:"not null"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Profile struct {
	ID               uint   `gorm:"primaryKey"`
	UserID           uint   `gorm:"unique;not null"` // Foreign key for User
	Bio              string
	ProfilePictureURL string
}

var db *gorm.DB
var err error

func main() {
	// Set up PostgreSQL connection info
	dsn := "host=localhost user=myuser password=password dbname=mydb5 port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Set up connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to set up connection pooling:", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(0)

	// Automatically migrate the User and Profile models
	err = db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database migration completed successfully.")

	// Insert a User and associated Profile in a single transaction
	insertUserWithProfile()

	// Query users with their profiles
	queryUsersWithProfiles()

	// Update a user's profile
	updateUserProfile(1, "Updated bio", "https://example.com/new-profile-picture.png")

	// Delete a user and associated profile
	deleteUser(1)
}


func insertUserWithProfile() {
	user := User{
		Name: "Alice",
		Age:  30,
		Profile: Profile{
			Bio:              "Hello, I am Alice!",
			ProfilePictureURL: "https://example.com/profile-picture.png",
		},
	}

	// Insert user and profile within a transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		fmt.Println("User and associated profile inserted successfully.")
		return nil
	})

	if err != nil {
		log.Fatal("Failed to insert user and profile:", err)
	}
}


func queryUsersWithProfiles() {
	var users []User
	result := db.Preload("Profile").Find(&users)
	if result.Error != nil {
		log.Fatal("Failed to query users and profiles:", result.Error)
	}

	fmt.Println("Users and their profiles:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Bio: %s, ProfilePictureURL: %s\n",
			user.ID, user.Name, user.Age, user.Profile.Bio, user.Profile.ProfilePictureURL)
	}
}



func updateUserProfile(userID uint, newBio, newProfilePictureURL string) {
	var profile Profile
	result := db.Model(&Profile{}).Where("user_id = ?", userID).First(&profile)
	if result.Error != nil {
		log.Fatal("Failed to find profile:", result.Error)
	}

	profile.Bio = newBio
	profile.ProfilePictureURL = newProfilePictureURL

	if err := db.Save(&profile).Error; err != nil {
		log.Fatal("Failed to update profile:", err)
	}
	fmt.Printf("Profile for user ID %d updated successfully.\n", userID)
}



func deleteUser(userID uint) {
	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		log.Fatal("Failed to find user:", result.Error)
	}

	if err := db.Delete(&user).Error; err != nil {
		log.Fatal("Failed to delete user and associated profile:", err)
	}
	fmt.Printf("User with ID %d and their associated profile deleted successfully.\n", userID)
}
