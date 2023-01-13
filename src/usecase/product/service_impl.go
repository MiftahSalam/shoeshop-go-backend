package product

import (
	"shoeshop-backend/src/interfaces/http/context"
)

func (s *service) GetProducts(ctx *context.ApplicationContext) (products []*ProductResponse, err error) {
	getProducts, err := s.pRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	products = s.toProductsResponse(getProducts)
	return
}
