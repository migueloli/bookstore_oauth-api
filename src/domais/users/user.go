package users

// User is the struct to process users.
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
}

// UserLoginRequest is the struct to login in the application.
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
