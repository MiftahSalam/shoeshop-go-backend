package order

import (
	"shoeshop-backend/src/interfaces/http/view/product"
	"shoeshop-backend/src/interfaces/http/view/user"
	"time"
)

type Item struct {
	Product   *product.Product `json:"product"`
	Name      string           `json:"name"`
	Quantity  int              `json:"quantity"`
	Price     float64          `json:"price"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

type ItemInput struct {
	ProductID string  `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderInput struct {
	Items           []*ItemInput   `json:"items"`
	ShippingAddress *ShippingInput `json:"shippingAddress"`
	PaymentMethod   string         `json:"paymentMethod"`
	TaxPrice        float64        `json:"taxPrice"`
	ShippingPrice   float64        `json:"shippingPrice"`
	TotalPrice      float64        `json:"totalPrice"`
}

type OrderResponse struct {
	ID              string         `json:"id"`
	User            *user.User     `json:"user"`
	Items           []*Item        `json:"items"`
	ShippingAddress *Shipping      `json:"shippingAddress"`
	PaymentMethod   string         `json:"paymentMethod"`
	PaymentStatus   *PaymentResult `json:"paymentStatus"`
	TaxPrice        float64        `json:"taxPrice"`
	ShippingPrice   float64        `json:"shippingPrice"`
	TotalPrice      float64        `json:"totalPrice"`
	IsPaid          bool           `json:"isPaid"`
	PaidAt          time.Time      `json:"paidAt"`
	IsDelivered     bool           `json:"isDelivered"`
	DeliveredAt     *time.Time     `json:"deliveredAt"`
	CreatedAt       *time.Time     `json:"createdAt"`
}

type PaymentResult struct {
	ID         string    `json:"id"`
	Status     string    `json:"status"`
	Email      string    `json:"email"`
	UpdateTime time.Time `json:"updateTime"`
}

type PaymentResultInput struct {
	ID         string    `json:"id"`
	Status     string    `json:"status"`
	Email      string    `json:"email"`
	UpdateTime time.Time `json:"updateTime"`
}

type Shipping struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type ShippingInput struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}
