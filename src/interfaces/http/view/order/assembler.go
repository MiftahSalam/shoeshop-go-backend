package order

import "shoeshop-backend/src/usecase/order"

func (oR *Order) ToOrderRequest() *order.OrderRequest {
	items := []*order.Item{}
	for _, item := range oR.Items {
		itemReq := &order.Item{
			Name:     item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
		}

		items = append(items, itemReq)
	}

	return &order.OrderRequest{
		Items:           items,
		ShippingAddress: order.Shipping(oR.ShippingAddress),
		PaymentMethod:   oR.PaymentMethod,
		TaxPrice:        oR.TaxPrice,
		ShippingPrice:   oR.ShippingPrice,
		TotalPrice:      oR.TotalPrice,
	}
}
