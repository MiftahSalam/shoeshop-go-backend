package order

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/usecase/order"
)

type (
	Service interface {
		CreateOrder(ctx *context.ApplicationContext, userId string, order *OrderInput) (*OrderResponse, error)
		GetOrder(ctx *context.ApplicationContext, orderId string) (*OrderResponse, error)
		GetOrders(ctx *context.ApplicationContext, userId string) ([]*OrderResponse, error)
		PayOrder(ctx *context.ApplicationContext, orderId string, payment PaymentResultInput) (*OrderResponse, error)
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

func (s *service) PayOrder(ctx *context.ApplicationContext, orderId string, payment PaymentResultInput) (out *OrderResponse, err error) {
	resp, err := s.oUC.PayOrder(ctx, orderId, payment.toOrderPaymentResult())
	if err != nil {
		return
	}

	out = toOrderResponse(resp)
	return
}

func (s *service) CreateOrder(ctx *context.ApplicationContext, userId string, order *OrderInput) (out *OrderResponse, err error) {
	resp, err := s.oUC.CreateOrder(ctx, userId, order.ToOrderRequest())
	if err != nil {
		return
	}

	out = toOrderResponse(resp)
	return
}

func (s *service) GetOrders(ctx *context.ApplicationContext, userId string) (out []*OrderResponse, err error) {
	resp, err := s.oUC.GetOrders(ctx, userId)
	if err != nil {
		return
	}

	out = toOrdersResponse(resp)
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
