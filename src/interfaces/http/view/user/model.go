package user

import "time"

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Token     string     `json:"token"`
	IsAdmin   bool       `json:"isAdmin"`
	CreatedAt *time.Time `json:"createdAt"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
