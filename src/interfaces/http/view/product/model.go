package product

import (
	"shoeshop-backend/src/interfaces/http/view/user"
	"shoeshop-backend/src/usecase/product"
)

type CreateProductRequest struct {
}

type CreateProductResponse struct {
}

type ProductResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CustomField1 string `json:"custom_field_1"`
}

type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  *string   `json:"description"`
	ImageURL     *string   `json:"imageUrl"`
	Rating       int       `json:"rating"`
	Price        float64   `json:"price"`
	NumReviews   int       `json:"numReviews"`
	CountInStock int       `json:"countInStock"`
	Reviews      []*Review `json:"reviews"`
}

type Review struct {
	Name    string     `json:"name"`
	Rating  int        `json:"rating"`
	Comment string     `json:"comment"`
	User    *user.User `json:"user"`
}

func ToProductsResponse(productsResp []*product.ProductResponse) []*Product {
	products := []*Product{}
	for _, v := range productsResp {
		products = append(products, toProductResponse(v))
	}

	return products
}

func (r *CreateProductRequest) ToProductUC() *product.ProductRequest {
	return &product.ProductRequest{}
}

func ToProductsResponseTest(productsResp []*product.ProductResponse) []*ProductResponse {
	products := []*ProductResponse{}
	for _, v := range productsResp {
		products = append(products, toProductResponseTest(v))
	}

	return products
}
