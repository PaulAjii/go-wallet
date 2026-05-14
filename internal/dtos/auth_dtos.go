package dtos

type RegisterRequest struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID         string `json:"id"`
	FullName   string `json:"fullName"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
