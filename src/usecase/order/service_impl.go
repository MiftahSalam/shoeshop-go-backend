package order

import (
	"fmt"
	"shoeshop-backend/src/domain/order"
	"shoeshop-backend/src/interfaces/http/context"
	"time"

	"github.com/google/uuid"
)

func (s *service) PayOrder(ctx *context.ApplicationContext, orderId string, payment PaymentResult) (resp *OrderResponse, err error) {
	order, err := s.oRepo.GetById(ctx, orderId)
	if err != nil {
		return nil, err
	}

	updatedOrder := s.createUpdatePayment(order, payment)
	err = s.oRepo.UpdateColumn(ctx, updatedOrder)
	if err != nil {
		return nil, err
	}

	resp = entityToOrderResponse(order)
	return
}

func (s *service) GetOrder(ctx *context.ApplicationContext, orderId string) (resp *OrderResponse, err error) {
	order, err := s.oRepo.GetById(ctx, orderId)
	if err != nil {
		return nil, err
	}

	resp = entityToOrderResponse(order)
	return
}

func (s *service) CreateOrder(ctx *context.ApplicationContext, userId string, orderInput *OrderRequest) (resp *OrderResponse, err error) {
	user, err := s.uRepo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	items, err := s.createItemsDomain(ctx, orderInput)
	if err != nil {
		return nil, err
	}

	createOrder := orderInput.ToOrderDomain(ctx, user, items)
	err = s.oRepo.Save(ctx, createOrder)
	if err != nil {
		return nil, err
	}

	resp = entityToOrderResponse(createOrder)
	return
}

func (s *service) Migrate() {
	s.oRepo.AutoMigrate()
}

func (s *service) createUpdatePayment(orderData *order.Order, payment PaymentResult) *order.Order {
	orderData.PaymentStatus.ID = payment.ID
	orderData.PaymentStatus.Status = payment.Status
	orderData.PaymentStatus.Email = payment.Email
	orderData.PaymentStatus.UpdateTime = payment.UpdateTime

	return &order.Order{
		ID:     orderData.ID,
		IsPaid: true,
		PaidAt: time.Now(),
		PaymentStatus: order.PaymentResult{
			ID:         payment.ID,
			Status:     payment.Status,
			Email:      payment.Email,
			UpdateTime: payment.UpdateTime,
		},
	}
}

func (s *service) createItemsDomain(ctx *context.ApplicationContext, orderInput *OrderRequest) ([]*order.Item, error) {
	items := []*order.Item{}
	for _, item := range orderInput.Items {
		productItem, err := s.pRepo.GetById(ctx, item.ProductId)
		if err != nil {
			ctx.Logger.Error(fmt.Sprintf("failed to create item domain with product item. error %v", err))
			return nil, err
		}

		itemDomain := &order.Item{
			ID:          uuid.New(),
			Name:        productItem.Name,
			Quantity:    item.Quantity,
			Price:       float64(item.Quantity) * productItem.Price,
			CreatedAt:   time.Now(),
			UpdatedDate: time.Now(),
			Product:     productItem,
		}

		items = append(items, itemDomain)
	}

	return items, nil
}
