package product

import "shoeshop-backend/src/usecase/product"

func toProductResponse(productResp *product.ProductResponse) *Product {
	return &Product{
		ID:           productResp.ID,
		Name:         productResp.Name,
		Description:  &productResp.Description,
		ImageURL:     &productResp.ImageUrl,
		Rating:       int(productResp.Rating),
		Price:        productResp.Price,
		NumReviews:   3, //temp
		CountInStock: int(productResp.StockCount),
		Reviews:      []*Review{}, //temp
	}
}

func toProductResponseTest(productResp *product.ProductResponse) *ProductResponse {
	return &ProductResponse{
		ID:           productResp.ID,
		Name:         productResp.Name,
		CustomField1: "Custom",
	}
}
