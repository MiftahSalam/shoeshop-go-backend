package order

import (
	"shoeshop-backend/src/domain/order"
)

type (
	Service interface {
		Migrate()
	}

	service struct {
		oRepo order.Repository
	}
)

func NewService(oRepo order.Repository) Service {
	if oRepo == nil {
		panic("please provide product repository")
	}

	return &service{
		oRepo: oRepo,
	}
}
