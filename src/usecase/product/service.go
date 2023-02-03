package product

import (
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/domain/user"
	"shoeshop-backend/src/interfaces/http/context"
)

type (
	Service interface {
		GetProducts(ctx *context.ApplicationContext, keyword string, page, limit int) (products []*ProductResponse, err error)
		GetProduct(ctx *context.ApplicationContext, id string) (product *ProductResponse, err error)
		CreateReview(ctx *context.ApplicationContext, userId string, review ReviewInput) (string, error)
		Migrate()
	}

	service struct {
		pRepo product.Repository
		uRepo user.Repository
	}
)

func NewService(pRepo product.Repository, uRepo user.Repository) Service {
	if pRepo == nil {
		panic("please provide product repository")
	}
	if uRepo == nil {
		panic("please provide user repository")
	}

	return &service{
		pRepo: pRepo,
		uRepo: uRepo,
	}
}
