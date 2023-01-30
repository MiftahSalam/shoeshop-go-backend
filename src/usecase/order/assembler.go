package order

import (
	"github.com/google/uuid"

	"shoeshop-backend/src/domain/order"
	"shoeshop-backend/src/domain/user"
	"shoeshop-backend/src/interfaces/http/context"
	uUC "shoeshop-backend/src/usecase/user"
)

func entityToOrderResponse(entity *order.Order) *OrderResponse {
	return &OrderResponse{
		ID: entity.ID.String(),
		User: &uUC.UserResponse{
			ID:        entity.User.ID.String(),
			Name:      entity.User.Name,
			Email:     entity.User.Email,
			IsAdmin:   entity.User.IsAdmin,
			CreatedAt: &entity.CreatedAt,
		},
		Items:           itemsDomainToItemsResponse(entity.Items),
		ShippingAddress: Shipping(entity.ShippingAddress),
		PaymentMethod:   entity.PaymentMethod,
		PaymentStatus:   PaymentResult(entity.PaymentStatus),
		TaxPrice:        entity.TaxPrice,
		ShippingPrice:   entity.ShippingPrice,
		TotalPrice:      entity.TotalPrice,
		IsPaid:          entity.IsPaid,
		PaidAt:          entity.PaidAt,
		IsDelivered:     entity.IsDelivered,
		DeliveredAt:     entity.DeliveredAt,
	}
}

func itemsDomainToItemsResponse(itemsInput []*order.Item) []*Item {
	items := []*Item{}
	for _, item := range itemsInput {
		itemResponse := &Item{
			ProductId: item.ProductId,
			Name:      item.Name,
			Quantity:  item.Quantity,
			Price:     item.Price,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedDate,
		}

		items = append(items, itemResponse)
	}

	return items
}
func (uR *OrderRequest) ToOrderDomain(ctx *context.ApplicationContext, user *user.User, items []*order.Item) *order.Order {
	return &order.Order{
		ID:              uuid.New(),
		User:            user,
		UserId:          user.ID.String(),
		Items:           items,
		ShippingAddress: order.Shipping(uR.ShippingAddress),
		PaymentMethod:   uR.PaymentMethod,
		PaymentStatus:   order.PaymentResult{},
		TaxPrice:        uR.TaxPrice,
		ShippingPrice:   uR.ShippingPrice, TotalPrice: uR.TotalPrice,
		IsPaid:      false,
		IsDelivered: false,
	}
}
