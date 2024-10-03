package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// PostgreSQL connection information
const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"     // replace with your PostgreSQL username
	password = "password" // replace with your PostgreSQL password
	dbname   = "mydb4"       // replace with your PostgreSQL database name
)

// Define the User struct
type User struct {
	Name string
	Age  int
}

func main() {
	// Create a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the database connection with connection pooling
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Set connection pooling settings
	db.SetMaxOpenConns(25)  // Max open connections
	db.SetMaxIdleConns(25)  // Max idle connections
	db.SetConnMaxLifetime(0) // No timeout for now

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
	fmt.Println("Successfully connected to the database!")

	// Create a users table with constraints
	createTable(db)

	// Insert multiple users using a transaction
	insertUsersTransaction(db, []User{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	})

	// Query users with filtering and pagination
	queryUsers(db, 25, 1, 2)

	// Update a user's details
	updateUser(db, 1, "Alicia", 31)

	// Delete a user by ID
	deleteUser(db, 3)
}

// Function to create a users table with constraints
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		age INT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	fmt.Println("Table created successfully.")
}

// Function to insert multiple users within a transaction
func insertUsersTransaction(db *sql.DB, users []User) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin transaction:", err)
	}

	query := `INSERT INTO users (name, age) VALUES ($1, $2)`
	for _, user := range users {
		_, err := tx.Exec(query, user.Name, user.Age)
		if err != nil {
			tx.Rollback() // Rollback on any error
			log.Fatal("Failed to insert user, rolling back:", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Failed to commit transaction:", err)
	}
	fmt.Println("All users inserted successfully.")
}

// Function to query users with filtering and pagination
func queryUsers(db *sql.DB, minAge int, page, pageSize int) {
	offset := (page - 1) * pageSize
	query := `SELECT id, name, age FROM users WHERE age >= $1 ORDER BY id LIMIT $2 OFFSET $3`
	rows, err := db.Query(query, minAge, pageSize, offset)
	if err != nil {
		log.Fatal("Failed to query users:", err)
	}
	defer rows.Close()

	fmt.Println("Queried Users:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("Failed to scan user:", err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}
}

// Function to update a user's details
func updateUser(db *sql.DB, id int, newName string, newAge int) {
	query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`
	res, err := db.Exec(query, newName, newAge, id)
	if err != nil {
		log.Fatal("Failed to update user:", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error fetching rows affected:", err)
	}

	if count == 0 {
		fmt.Printf("No user found with ID %d to update.\n", id)
	} else {
		fmt.Printf("User with ID %d updated successfully.\n", id)
	}
}

// Function to delete a user by ID
func deleteUser(db *sql.DB, id int) {
	query := `DELETE FROM users WHERE id = $1`
	res, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("Failed to delete user:", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error fetching rows affected:", err)
	}

	if count == 0 {
		fmt.Printf("No user found with ID %d to delete.\n", id)
	} else {
		fmt.Printf("User with ID %d deleted successfully.\n", id)
	}
}
