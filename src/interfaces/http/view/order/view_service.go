package order

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/usecase/order"
)

type (
	Service interface {
		CreateOrder(ctx *context.ApplicationContext, userId string, order *Order) (*order.OrderResponse, error)
	}

	service struct {
		oUC order.Service
	}
)

func NewService(oUC order.Service) Service {
	if oUC == nil {
		panic("order usecase is nil")
	}

	return &service{oUC: oUC}
}

func (s *service) CreateOrder(ctx *context.ApplicationContext, userId string, order *Order) (out *order.OrderResponse, err error) {
	out, err = s.oUC.CreateOrder(ctx, userId, order.ToOrderRequest())
	if err != nil {
		return
	}

	return
}
