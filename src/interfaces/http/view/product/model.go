package product

import "shoeshop-backend/src/usecase/product"

type CreateProductRequest struct {
}

type CreateProductResponse struct {
}

type ProductResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CustomField1 string `json:"custom_field_1"`
}

func ToProductsResponse(productsResp []*product.ProductResponse) []*ProductResponse {
	products := []*ProductResponse{}
	for _, v := range productsResp {
		products = append(products, toProductResponse(v))
	}

	return products
}

func (r *CreateProductRequest) ToProductUC() *product.ProductRequest {
	return &product.ProductRequest{}
}
