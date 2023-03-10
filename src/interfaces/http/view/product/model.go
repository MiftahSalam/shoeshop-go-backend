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

type Products struct {
	ProductList []*Product `json:"product_list"`
	TotalData   int64      `json:"total_data"`
}

type Search struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}

type ReviewInput struct {
	ProductID string `json:"productId"`
	Rating    int    `json:"Rating"`
	Comment   string `json:"Comment"`
}

type Review struct {
	ID          string     `json:"id"`
	Rating      int        `json:"rating"`
	Comment     string     `json:"comment"`
	User        *user.User `json:"user"`
	CreatedDate string     `json:"created_date"`
}

func ToProductsResponse(productsResp *product.ProductsResponse) *Products {
	products := []*Product{}
	for _, v := range productsResp.Products {
		products = append(products, toProductResponse(v))
	}

	return &Products{
		ProductList: products,
		TotalData:   productsResp.TotalData,
	}
}

func (r *CreateProductRequest) ToProductUC() *product.ProductRequest {
	return &product.ProductRequest{}
}

func (s *Search) validate() (keyword string, page int, limit int) {
	keyword, page, limit = s.Keyword, s.Page, s.Limit
	if s.Page < 1 {
		page = 1
	} else if s.Page > 1000 {
		page = 1000
	}

	if s.Limit < 1 {
		limit = 10
	} else if s.Limit > 100 {
		limit = 100
	}

	return
}

func ToProductsResponseTest(productsResp *product.ProductsResponse) []*ProductResponse {
	products := []*ProductResponse{}
	for _, v := range productsResp.Products {
		products = append(products, toProductResponseTest(v))
	}

	return products
}
