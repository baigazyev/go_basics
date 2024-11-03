package user

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // Add Role field
}

// User storage (simulated database)
var users = make(map[int]User)
var nextID = 1

func AddUser(username, password, role string) User {
	user := User{
		ID:       nextID,
		Username: username,
		Password: password,
		Role:     role, // Assign role
	}
	users[nextID] = user
	nextID++
	return user
}

// GetUser returns a user by ID
func GetUser(id int) (User, bool) {
	user, exists := users[id]
	return user, exists
}

// GetAllUsers returns all users
func GetAllUsers() []User {
	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}
	return userList
}

// UpdateUser updates user details by ID
func UpdateUser(id int, username, password string) (User, bool) {
	user, exists := users[id]
	if exists {
		user.Username = username
		if password != "" {
			// Only update password if it's provided
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err == nil {
				user.Password = string(hashedPassword)
			}
		}
		users[id] = user
	}
	return user, exists
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) bool {
	_, exists := users[id]
	if exists {
		delete(users, id)
	}
	return exists
}

// GetUserByUsername retrieves a user by username for authentication
func GetUserByUsername(username string) (User, bool) {
	for _, user := range users {
		if user.Username == username {
			return user, true
		}
	}
	return User{}, false
}
