package user

import (
	"time"
)

type UserResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	IsAdmin   bool       `json:"isAdmin"`
	CreatedAt *time.Time `json:"createdAt"`
}

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UserUpdate struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
