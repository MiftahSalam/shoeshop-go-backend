package order

import (
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/interfaces/http/view/user"
	"time"
)

type OrderResponse struct {
	ID              string        `json:"id"`
	User            *user.User    `json:"user"`
	Items           []*Item       `json:"items"`
	ShippingAddress Shipping      `json:"shipping_address"`
	PaymentMethod   string        `json:"payment_method"`
	PaymentStatus   PaymentResult `json:"payment_status"`
	TaxPrice        float64       `json:"tax_price"`
	ShippingPrice   float64       `json:"shipping_price"`
	TotalPrice      float64       `json:"total_price"`
	IsPaid          bool          `json:"is_paid"`
	PaidAt          time.Time     `json:"paid_at"`
	IsDelivered     bool          `json:"is_delivered"`
	DeliveredAt     time.Time     `json:"delivered_at"`
	CreatedAt       time.Time     `json:"created_at"`
}

type OrderRequest struct {
	Items           []*Item  `json:"items"`
	ShippingAddress Shipping `json:"shipping_address"`
	PaymentMethod   string   `json:"payment_method"`
	TaxPrice        float64  `json:"tax_price"`
	ShippingPrice   float64  `json:"shipping_price"`
	TotalPrice      float64  `json:"total_price"`
}

type Item struct {
	Product   *product.Product `json:"product"`
	ProductId string           `json:"product_id"`
	Name      string           `json:"name"`
	Quantity  int              `json:"quantity"`
	Price     float64          `json:"price"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type Shipping struct {
	Address    string `json:"name"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type PaymentResult struct {
	Status     string    `json:"status"`
	UpdateTime time.Time `json:"update_time"`
	Email      string    `json:"email"`
}
