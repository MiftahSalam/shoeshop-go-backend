package user

import (
	"shoeshop-backend/src/interfaces/http/context"
)

type Repository interface {
	AutoMigrate()
	Save(ctx *context.ApplicationContext, userInput *User) (err error)
	GetByEmail(ctx *context.ApplicationContext, email string) (user *User, err error)
	GetById(ctx *context.ApplicationContext, id string) (user *User, err error)
}
