package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    _ "github.com/lib/pq"
)

// Database variables
var (
    db    *sql.DB
    gormDB *gorm.DB
)

// User struct
type User struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name" gorm:"unique;not null"`
    Age   int    `json:"age" gorm:"not null"`
}

func main() {
    // PostgreSQL connection setup for SQL and GORM
    var err error
    db, err = sql.Open("postgres", "user=myuser password=password dbname=mydb8 sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to the database using sql:", err)
    }
    defer db.Close()

    dsn := "host=localhost user=myuser password=password dbname=mydb8 port=5432 sslmode=disable TimeZone=Asia/Almaty"
    gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database using GORM:", err)
    }

    // Setup connection pooling
    sqlDB, err := gormDB.DB()
    if err != nil {
        log.Fatal(err)
    }
    sqlDB.SetMaxOpenConns(10)
    sqlDB.SetMaxIdleConns(5)

    // Auto migrate the user table
    gormDB.AutoMigrate(&User{})

    // Setup router
    router := mux.NewRouter()

    // SQL-based routes
    router.HandleFunc("/sql/users", getUsersSQL).Methods("GET")
    router.HandleFunc("/sql/users", createUserSQL).Methods("POST")
    router.HandleFunc("/sql/users/{id}", updateUserSQL).Methods("PUT")
    router.HandleFunc("/sql/users/{id}", deleteUserSQL).Methods("DELETE")

    // GORM-based routes
    router.HandleFunc("/gorm/users", getGormUsers).Methods("GET")
    router.HandleFunc("/gorm/users", createGormUser).Methods("POST")
    router.HandleFunc("/gorm/users/{id}", updateGormUser).Methods("PUT")
    router.HandleFunc("/gorm/users/{id}", deleteGormUser).Methods("DELETE")

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

//////////////////////// SQL Handlers ////////////////////////

func getUsersSQL(w http.ResponseWriter, r *http.Request) {
    ageFilter := r.URL.Query().Get("age")
    sort := r.URL.Query().Get("sort")
    limit := r.URL.Query().Get("limit")
    offset := r.URL.Query().Get("offset")

    query := "SELECT id, name, age FROM users WHERE 1=1"
    params := []interface{}{}
    if ageFilter != "" {
        query += " AND age >= $1"
        params = append(params, ageFilter)
    }
    if sort == "asc" {
        query += " ORDER BY name ASC"
    } else if sort == "desc" {
        query += " ORDER BY name DESC"
    }
    query += " LIMIT $2 OFFSET $3"
    params = append(params, limit, offset)

    rows, err := db.Query(query, params...)
    if err != nil {
        http.Error(w, "Error querying users: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
            http.Error(w, "Error scanning user: "+err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createUserSQL(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if user.Name == "" {
        http.Error(w, "Name cannot be empty", http.StatusBadRequest)
        return
    }

    // Check if name is unique
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM users WHERE name=$1", user.Name).Scan(&count)
    if err != nil {
        http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    if count > 0 {
        http.Error(w, "Name already exists", http.StatusConflict)
        return
    }

    _, err = db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
    if err != nil {
        http.Error(w, "Error inserting user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func updateUserSQL(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", user.Name, user.Age, id)
    if err != nil {
        http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := db.Exec("DELETE FROM users WHERE id=$1", id)
    if err != nil {
        http.Error(w, "Error deleting user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

//////////////////////// GORM Handlers ////////////////////////

func getGormUsers(w http.ResponseWriter, r *http.Request) {
    ageFilter := r.URL.Query().Get("age")
    sort := r.URL.Query().Get("sort")
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

    var users []User
    query := gormDB.Limit(limit).Offset(offset)

    if ageFilter != "" {
        query = query.Where("age >= ?", ageFilter)
    }
    if sort == "asc" {
        query = query.Order("name asc")
    } else if sort == "desc" {
        query = query.Order("name desc")
    }

    result := query.Find(&users)
    if result.Error != nil {
        http.Error(w, "Error querying users: "+result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createGormUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := gormDB.Create(&user).Error; err != nil {
        http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func updateGormUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := gormDB.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
        http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func deleteGormUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    if err := gormDB.Delete(&User{}, id).Error; err != nil {
        http.Error(w, "Error deleting user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
