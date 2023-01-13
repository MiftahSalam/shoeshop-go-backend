package product

import (
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/interfaces/http/context"
)

type (
	Service interface {
		GetProducts(ctx *context.ApplicationContext) (products []*ProductResponse, err error)
	}

	service struct {
		pRepo product.Repository
	}
)

func NewService(pRepo product.Repository) Service {
	if pRepo == nil {
		panic("please provide product repository")
	}

	return &service{
		pRepo: pRepo,
	}
}
