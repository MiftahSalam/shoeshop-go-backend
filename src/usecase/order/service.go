package order

import (
	"shoeshop-backend/src/domain/order"
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/domain/user"
	"shoeshop-backend/src/interfaces/http/context"
)

type (
	Service interface {
		Migrate()
		CreateOrder(ctx *context.ApplicationContext, userId string, orderInput *OrderRequest) (resp *OrderResponse, err error)
	}

	service struct {
		oRepo order.Repository
		pRepo product.Repository
		uRepo user.Repository
	}
)

func NewService(oRepo order.Repository, pRepo product.Repository, uRepo user.Repository) Service {
	if oRepo == nil {
		panic("please provide order repository")
	}

	if pRepo == nil {
		panic("please provide product repository")
	}

	if uRepo == nil {
		panic("please provide user repository")
	}

	return &service{
		oRepo: oRepo,
		pRepo: pRepo,
		uRepo: uRepo,
	}
}
