package product

import (
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/domain/user"
	"time"

	"github.com/google/uuid"
)

func (rIn ReviewInput) ToReviewDomain(userIn *user.User, productIn *product.Product) *product.Review {
	return &product.Review{
		ID:          uuid.New(),
		Rating:      rIn.Rating,
		Comment:     rIn.Comment,
		Product:     productIn,
		User:        userIn,
		CreatedDate: time.Now(),
	}
}

func (s *service) toProductsResponse(entities []*product.Product) []*ProductResponse {
	products := []*ProductResponse{}
	for _, v := range entities {
		products = append(products, entityToProdcutResponse(v))
	}

	return products
}

func entityToProdcutResponse(entity *product.Product) *ProductResponse {
	return &ProductResponse{
		ID:          entity.ID.String(),
		Name:        entity.Name,
		Description: entity.Description,
		ImageUrl:    entity.ImageUrl,
		Rating:      entity.Rating,
		Price:       entity.Price,
		StockCount:  entity.StockCount,
		CreatedDate: entity.CreatedDate.String(),
		Reviews:     entity.Reviews,
	}
}
