package user

import (
	"shoeshop-backend/src/domain/user"
	"shoeshop-backend/src/interfaces/http/context"
)

type (
	Service interface {
		Migrate()
		GetByEmail(ctx *context.ApplicationContext, email string) (resp *UserResponse, err error)
		GetById(ctx *context.ApplicationContext, id string) (resp *UserResponse, err error)
		LoginUser(ctx *context.ApplicationContext, email, password string) (resp *UserResponse, err error)
	}

	service struct {
		uRepo user.Repository
	}
)

func NewService(uRepo user.Repository) Service {
	if uRepo == nil {
		panic("please provide user repository")
	}

	return &service{
		uRepo: uRepo,
	}
}
