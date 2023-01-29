package user

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/usecase/user"
)

type (
	Service interface {
		GetUser(ctx *context.ApplicationContext, email string) (*User, error)
		LoginUser(ctx *context.ApplicationContext, email, password string) (*User, error)
	}

	service struct {
		uUC user.Service
	}
)

func NewService(uUC user.Service) Service {
	if uUC == nil {
		panic("product usecase is nil")
	}

	return &service{uUC: uUC}
}

func (s *service) LoginUser(ctx *context.ApplicationContext, email, password string) (out *User, err error) {
	res, err := s.uUC.LoginUser(ctx, email, password)
	if err != nil {
		return
	}

	out = toUserResponse(res)

	return
}

func (s *service) GetUser(ctx *context.ApplicationContext, email string) (out *User, err error) {
	res, err := s.uUC.GetByEmail(ctx, email)
	if err != nil {
		return
	}

	out = toUserResponse(res)

	return
}
