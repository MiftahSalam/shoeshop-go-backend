package product

import "shoeshop-backend/src/usecase/product"

func toProductResponse(productResp *product.ProductResponse) *ProductResponse {
	return &ProductResponse{
		ID:           productResp.ID,
		Name:         productResp.Name,
		CustomField1: "Custom",
	}
}
