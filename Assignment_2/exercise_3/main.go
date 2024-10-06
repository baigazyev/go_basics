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

var (
    sqlDB  *sql.DB
    gormDB *gorm.DB
)

// User represents the user model
type User struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

func main() {
    // Connect to SQL database
    var err error
    sqlDB, err = sql.Open("postgres", "user=myuser password=password dbname=mydb7 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer sqlDB.Close()

    // Connect to GORM database
    dsn := "host=localhost user=myuser password=password dbname=mydb7 port=5432 sslmode=disable TimeZone=Asia/Almaty"
    gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Migrate schema
    gormDB.AutoMigrate(&User{})

    router := mux.NewRouter()

    // SQL-based handlers
    router.HandleFunc("/sql/users", getUsersSQL).Methods("GET")
    router.HandleFunc("/sql/user", createUserSQL).Methods("POST")
    router.HandleFunc("/sql/user/{id}", updateUserSQL).Methods("PUT")
    router.HandleFunc("/sql/user/{id}", deleteUserSQL).Methods("DELETE")

    // GORM-based handlers
    router.HandleFunc("/gorm/users", getGormUsers).Methods("GET")
    router.HandleFunc("/gorm/user", createGormUser).Methods("POST")
    router.HandleFunc("/gorm/user/{id}", updateGormUser).Methods("PUT")
    router.HandleFunc("/gorm/user/{id}", deleteGormUser).Methods("DELETE")

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

/////////////////////// SQL Handlers ///////////////////////

func getUsersSQL(w http.ResponseWriter, r *http.Request) {
    rows, err := sqlDB.Query("SELECT id, name, email, age FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    users := []User{}
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
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
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := sqlDB.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", user.Name, user.Email, user.Age)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func updateUserSQL(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = sqlDB.Exec("UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4", user.Name, user.Email, user.Age, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    _, err = sqlDB.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

/////////////////////// GORM Handlers ///////////////////////

func getGormUsers(w http.ResponseWriter, r *http.Request) {
    var users []User
    result := gormDB.Find(&users)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createGormUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := gormDB.Create(&user)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func updateGormUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := gormDB.Model(&User{}).Where("id = ?", id).Updates(user)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func deleteGormUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    result := gormDB.Delete(&User{}, id)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
