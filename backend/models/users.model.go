package models

type User struct {
	ID           string `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password" db:"password"`
	Role         string `json:"role" db:"role"`
}
