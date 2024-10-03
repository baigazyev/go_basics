package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQL connection information
const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "password"
	dbname   = "mydb6"
)

// User model for GORM
type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
	Age  int    `gorm:"not null" json:"age"`
}

var (
	dbSQL  *sql.DB
	dbGORM *gorm.DB
	err    error
)

func main() {
	// Set up PostgreSQL connection for database/sql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbSQL, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer dbSQL.Close()

	// Set up PostgreSQL connection for GORM
	dbGORM, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database using GORM:", err)
	}

	// Auto-migrate the User table
	err = dbGORM.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setting up routes
	router := mux.NewRouter()

	// Routes using database/sql
	router.HandleFunc("/users/sql", getUsersSQL).Methods("GET")
	router.HandleFunc("/users/sql", createUserSQL).Methods("POST")
	router.HandleFunc("/users/sql/{id:[0-9]+}", updateUserSQL).Methods("PUT")
	router.HandleFunc("/users/sql/{id:[0-9]+}", deleteUserSQL).Methods("DELETE")

	// Routes using GORM
	router.HandleFunc("/users/gorm", getUsersGORM).Methods("GET")
	router.HandleFunc("/users/gorm", createUserGORM).Methods("POST")
	router.HandleFunc("/users/gorm/{id:[0-9]+}", updateUserGORM).Methods("PUT")
	router.HandleFunc("/users/gorm/{id:[0-9]+}", deleteUserGORM).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Direct SQL Handlers

// Fetch all users with filtering, sorting, and pagination using database/sql
func getUsersSQL(w http.ResponseWriter, r *http.Request) {
	ageFilter := r.URL.Query().Get("age")
	sortBy := r.URL.Query().Get("sort")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := `SELECT id, name, age FROM users WHERE ($1::int IS NULL OR age >= $1)`
	if sortBy == "name" {
		query += ` ORDER BY name`
	} else {
		query += ` ORDER BY id`
	}
	query += ` LIMIT $2 OFFSET $3`

	var rows *sql.Rows
	var err error
	if ageFilter != "" {
		age, _ := strconv.Atoi(ageFilter)
		rows, err = dbSQL.Query(query, age, pageSize, offset)
	} else {
		rows, err = dbSQL.Query(query, nil, pageSize, offset)
	}

	if err != nil {
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// Insert a new user using database/sql
func createUserSQL(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	err := dbSQL.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Failed to insert user. Ensure the name is unique.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Update a user by ID using database/sql
func updateUserSQL(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`
	_, err := dbSQL.Exec(query, user.Name, user.Age, id)
	if err != nil {
		http.Error(w, "Failed to update user. Ensure the name is unique.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

// Delete a user by ID using database/sql
func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	query := `DELETE FROM users WHERE id = $1`
	_, err := dbSQL.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

// GORM Handlers

// Fetch all users with filtering, sorting, and pagination using GORM
func getUsersGORM(w http.ResponseWriter, r *http.Request) {
	ageFilter := r.URL.Query().Get("age")
	sortBy := r.URL.Query().Get("sort")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var users []User
	query := dbGORM.Model(&User{})
	if ageFilter != "" {
		age, _ := strconv.Atoi(ageFilter)
		query = query.Where("age >= ?", age)
	}

	if sortBy == "name" {
		query = query.Order("name")
	} else {
		query = query.Order("id")
	}

	result := query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)
	if result.Error != nil {
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// Insert a new user using GORM
func createUserGORM(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := dbGORM.Create(&user).Error; err != nil {
		http.Error(w, "Failed to insert user. Ensure the name is unique.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Update a user by ID using GORM
func updateUserGORM(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := dbGORM.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		http.Error(w, "Failed to update user. Ensure the name is unique.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

// Delete a user by ID using GORM
func deleteUserGORM(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := dbGORM.Delete(&User{}, id).Error; err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
