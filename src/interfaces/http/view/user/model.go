package user

import "time"

type User struct {
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	IsAdmin   bool       `json:"isAdmin"`
	CreatedAt *time.Time `json:"createdAt"`
}
