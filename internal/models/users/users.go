package users

import "github.com/PaulAjii/go-wallet/internal/models"

type UsersModel struct {
	models.BaseModel

	FullName   string `json:"fullName" db:"full_name"`
	Username   string `json:"username" db:"username"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"-" db:"password_hash"`
	IsVerified bool   `json:"isVerified" db:"is_verified"`
}
