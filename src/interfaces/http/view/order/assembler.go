package order

import (
	"shoeshop-backend/src/interfaces/http/view/product"
	oUC "shoeshop-backend/src/usecase/order"
)

func (oR *OrderInput) ToOrderRequest() *oUC.OrderRequest {
	items := []*oUC.Item{}
	for _, item := range oR.Items {
		itemReq := &oUC.Item{
			Quantity:  item.Quantity,
			Price:     item.Price,
			ProductId: item.ProductID,
		}

		items = append(items, itemReq)
	}

	return &oUC.OrderRequest{
		Items: items,
		ShippingAddress: oUC.Shipping{
			Address:    oR.ShippingAddress.Address,
			City:       oR.ShippingAddress.City,
			PostalCode: oR.ShippingAddress.PostalCode,
			Country:    oR.ShippingAddress.Country,
		},
		PaymentMethod: oR.PaymentMethod,
		TaxPrice:      oR.TaxPrice,
		ShippingPrice: oR.ShippingPrice,
		TotalPrice:    oR.TotalPrice,
	}
}

func toOrderResponse(order *oUC.OrderResponse) *OrderResponse {
	items := []*Item{}
	for _, item := range order.Items {
		itemReq := &Item{
			Name:     item.Product.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
			Product: &product.Product{
				ID:           item.Product.ID.String(),
				Name:         item.Product.Name,
				Description:  &item.Product.Description,
				ImageURL:     &item.Product.ImageUrl,
				Rating:       int(item.Product.Rating),
				Price:        item.Product.Price,
				CountInStock: int(item.Product.StockCount),
				Reviews:      []*product.Review{}, //temporary
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		items = append(items, itemReq)
	}

	return &OrderResponse{
		ID:    order.ID,
		User:  order.User,
		Items: items,
		ShippingAddress: &Shipping{
			Address:    order.ShippingAddress.Address,
			City:       order.ShippingAddress.City,
			PostalCode: order.ShippingAddress.PostalCode,
			Country:    order.ShippingAddress.Country,
		},
		PaymentMethod: order.PaymentMethod,
		PaymentStatus: &PaymentResult{
			Status:     order.PaymentStatus.Status,
			Email:      order.PaymentStatus.Email,
			UpdateTime: order.PaymentStatus.UpdateTime,
		},
		TaxPrice:      order.TaxPrice,
		ShippingPrice: order.ShippingPrice,
		TotalPrice:    order.TotalPrice,
		IsPaid:        order.IsPaid,
		PaidAt:        order.PaidAt,
		IsDelivered:   order.IsDelivered,
		DeliveredAt:   &order.DeliveredAt,
		CreatedAt:     &order.CreatedAt,
	}
}
