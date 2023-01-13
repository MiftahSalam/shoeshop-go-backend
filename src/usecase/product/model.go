package product

type ProductRequest struct {
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Rating      int8    `json:"rating"`
	Price       float64 `json:"price"`
	StockCount  float64 `json:"stock"`
	CreatedDate string  `json:"created_date"`
	DeletedDate string  `json:"deleted_date"`
	UpdatedDate string  `json:"updated_date"`
}
