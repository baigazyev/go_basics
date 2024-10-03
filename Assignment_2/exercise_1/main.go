package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// PostgreSQL connection info
const (
	host     = "localhost"
	port     = 5432        // default PostgreSQL port
	user     = "myuser" // replace with your PostgreSQL username
	password = "password" // replace with your PostgreSQL password
	dbname   = "mydb"   // replace with your PostgreSQL database name
)

func main() {
	// Create a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	fmt.Println("Successfully connected!")

	// Create the users table
	createTable(db)

	// Insert some records
	insertUser(db, "Alice", 30)
	insertUser(db, "Bob", 25)
	insertUser(db, "Charlie", 35)

	// Query and print all users
	queryUsers(db)
}

// Function to create a table
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	fmt.Println("Table created successfully.")
}

// Function to insert a user into the users table
func insertUser(db *sql.DB, name string, age int) {
	query := `INSERT INTO users (name, age) VALUES ($1, $2)`
	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Fatal("Failed to insert user:", err)
	}
	fmt.Printf("Inserted user: %s, age: %d\n", name, age)
}

// Function to query and print all users
func queryUsers(db *sql.DB) {
	query := `SELECT id, name, age FROM users`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Failed to query users:", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
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

	// Check for errors from iteration
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}


