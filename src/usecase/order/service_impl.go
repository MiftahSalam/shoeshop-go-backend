package order

import (
	"fmt"
	"shoeshop-backend/src/domain/order"
	"shoeshop-backend/src/interfaces/http/context"
	"time"

	"github.com/google/uuid"
)

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
