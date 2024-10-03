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

// PostgreSQL connection info
const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "password"
	dbname   = "mydb3"
)

// User model for GORM
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Age  int    `gorm:"not null"`
}

var (
	dbSQL  *sql.DB
	dbGORM *gorm.DB
	err    error
)

func main() {
	// Connect using database/sql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbSQL, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer dbSQL.Close()

	err = dbSQL.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database using database/sql:", err)
	}

	// Connect using GORM
	dsn := psqlInfo
	dbGORM, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database using GORM:", err)
	}

	// AutoMigrate User table using GORM
	err = dbGORM.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setting up routes
	router := mux.NewRouter()

	// Routes using database/sql
	router.HandleFunc("/users", getUsersSQL).Methods("GET")
	router.HandleFunc("/user", createUserSQL).Methods("POST")
	router.HandleFunc("/user/{id}", updateUserSQL).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUserSQL).Methods("DELETE")

	// Routes using GORM
	router.HandleFunc("/gorm/users", getUsersGORM).Methods("GET")
	router.HandleFunc("/gorm/user", createUserGORM).Methods("POST")
	router.HandleFunc("/gorm/user/{id}", updateUserGORM).Methods("PUT")
	router.HandleFunc("/gorm/user/{id}", deleteUserGORM).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}


func getUsersSQL(w http.ResponseWriter, r *http.Request) {
	rows, err := dbSQL.Query("SELECT id, name, age FROM users")
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


func createUserSQL(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	err := dbSQL.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}


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
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}



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


func getUsersGORM(w http.ResponseWriter, r *http.Request) {
	var users []User
	result := dbGORM.Find(&users)
	if result.Error != nil {
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}


func createUserGORM(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := dbGORM.Create(&user)
	if result.Error != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}


func updateUserGORM(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := dbGORM.Model(&User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}


func deleteUserGORM(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	result := dbGORM.Delete(&User{}, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
