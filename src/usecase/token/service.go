package token

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/config"
)

type (
	Service interface {
		Generate(ctx *context.ApplicationContext, userId string) (result string, err error)
	}

	service struct {
		cfg *config.Application
	}
)

func NewService(cfg *config.Application) Service {
	if cfg == nil {
		panic("please provide config")
	}
	return &service{cfg: cfg}
}
