package order

type Order struct {
	Items           []*Item  `json:"items"`
	ShippingAddress Shipping `json:"shipping_address"`
	PaymentMethod   string   `json:"payment_method"`
	TaxPrice        float64  `json:"tax_price"`
	ShippingPrice   float64  `json:"shipping_price"`
	TotalPrice      float64  `json:"total_price"`
}

type Item struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Shipping struct {
	Address    string `json:"name"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}
