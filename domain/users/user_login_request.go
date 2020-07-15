package users

// LoginRequest - Login request structure
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
