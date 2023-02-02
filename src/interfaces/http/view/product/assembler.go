package product

import (
	"shoeshop-backend/src/interfaces/http/view/user"
	"shoeshop-backend/src/usecase/product"
	uUC "shoeshop-backend/src/usecase/user"
)

func toProductResponse(productResp *product.ProductResponse) *Product {
	reviews := []*Review{}
	for _, review := range productResp.Reviews {
		reviewOut := &Review{
			Rating:  review.Rating,
			Comment: review.Comment,
			User:    user.ToUserResponse(uUC.EntityToUserResponse(review.User)),
		}

		reviews = append(reviews, reviewOut)
	}

	return &Product{
		ID:           productResp.ID,
		Name:         productResp.Name,
		Description:  &productResp.Description,
		ImageURL:     &productResp.ImageUrl,
		Rating:       int(productResp.Rating),
		Price:        productResp.Price,
		NumReviews:   len(productResp.Reviews), //temp
		CountInStock: int(productResp.StockCount),
		Reviews:      reviews,
	}
}

func toProductResponseTest(productResp *product.ProductResponse) *ProductResponse {
	return &ProductResponse{
		ID:           productResp.ID,
		Name:         productResp.Name,
		CustomField1: "Custom",
	}
}
