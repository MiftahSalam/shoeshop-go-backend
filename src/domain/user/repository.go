package user

import (
	"shoeshop-backend/src/interfaces/http/context"
)

type Repository interface {
	AutoMigrate()
	GetByEmail(ctx *context.ApplicationContext, email string) (user *User, err error)
	GetById(ctx *context.ApplicationContext, id string) (user *User, err error)
}
