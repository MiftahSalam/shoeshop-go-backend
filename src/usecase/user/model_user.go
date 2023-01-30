package user

import (
	"shoeshop-backend/src/domain/user"
	"shoeshop-backend/src/interfaces/http/context"
	"time"

	"github.com/google/uuid"
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

func (uR UserRegister) ToUserDomain(ctx *context.ApplicationContext) *user.User {
	hashedPassword, err := HashPassword(uR.Password)
	if err != nil {
		ctx.Logger.Error("failed create user domain. error hash password: " + err.Error())
		return nil
	}

	return &user.User{
		ID:          uuid.New(),
		Name:        uR.Name,
		Email:       uR.Email,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		Password:    hashedPassword,
	}
}
