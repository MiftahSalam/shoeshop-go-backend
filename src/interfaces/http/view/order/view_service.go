package order

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/usecase/order"
)

type (
	Service interface {
		CreateOrder(ctx *context.ApplicationContext, userId string, order *OrderInput) (*OrderResponse, error)
		GetOrder(ctx *context.ApplicationContext, orderId string) (*OrderResponse, error)
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

func (s *service) CreateOrder(ctx *context.ApplicationContext, userId string, order *OrderInput) (out *OrderResponse, err error) {
	resp, err := s.oUC.CreateOrder(ctx, userId, order.ToOrderRequest())
	if err != nil {
		return
	}

	out = toOrderResponse(resp)
	return
}

func (s *service) GetOrder(ctx *context.ApplicationContext, orderId string) (out *OrderResponse, err error) {
	resp, err := s.oUC.GetOrder(ctx, orderId)
	if err != nil {
		return
	}

	out = toOrderResponse(resp)
	return
}
