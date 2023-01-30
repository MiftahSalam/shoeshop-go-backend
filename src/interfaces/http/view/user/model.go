package user

import (
	"shoeshop-backend/src/usecase/user"
	"time"
)

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

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateProfile struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r Register) ToUserRegister() user.UserRegister {
	return user.UserRegister{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

func (uP UpdateProfile) ToUserUpdate(id string) user.UserUpdate {
	return user.UserUpdate{
		ID:       id,
		Name:     uP.Name,
		Email:    uP.Email,
		Password: uP.Password,
	}
}
