package product

import (
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/interfaces/http/context"
)

type (
	Service interface {
		GetProducts(ctx *context.ApplicationContext) (products []*ProductResponse, err error)
		GetProduct(ctx *context.ApplicationContext, id string) (product *ProductResponse, err error)
		Migrate()
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
