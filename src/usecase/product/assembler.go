package product

import "shoeshop-backend/src/domain/product"

func (s *service) toProductsResponse(entities []*product.Product) []*ProductResponse {
	products := []*ProductResponse{}
	for _, v := range entities {
		products = append(products, entityToProdcutResponse(v))
	}

	return products
}

func entityToProdcutResponse(entity *product.Product) *ProductResponse {
	return &ProductResponse{
		ID:   entity.ID.String(),
		Name: entity.Name,
	}
}
