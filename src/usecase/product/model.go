package product

import "shoeshop-backend/src/domain/product"

type ProductRequest struct {
}

type ProductsResponse struct {
	Products  []*ProductResponse `json:"products"`
	TotalData int64              `json:"total_data"`
}

type ProductResponse struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	ImageUrl    string            `json:"image_url"`
	Rating      int8              `json:"rating"`
	Price       float64           `json:"price"`
	StockCount  float64           `json:"stock"`
	Reviews     []*product.Review `json:"reviews"`
	CreatedDate string            `json:"created_date"`
	DeletedDate string            `json:"deleted_date"`
	UpdatedDate string            `json:"updated_date"`
}

type ReviewInput struct {
	ProductID string `json:"productId"`
	Rating    int    `json:"Rating"`
	Comment   string `json:"Comment"`
}
