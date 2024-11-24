package handlers

// "github.com/go-playground/validator/v10"

// var validate = validator.New()

// type LoginRequest struct {
// 	Username string `json:"username" validate:"required,min=3,max=50"`
// 	Password string `json:"password" validate:"required,min=8"`
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	var loginReq LoginRequest
// 	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
// 		http.Error(w, "Invalid request format", http.StatusBadRequest)
// 		return
// 	}

// 	if err := validate.Struct(loginReq); err != nil {
// 		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Continue with authentication...
// }
