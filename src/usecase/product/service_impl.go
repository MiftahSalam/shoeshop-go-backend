package product

import (
	"shoeshop-backend/src/interfaces/http/context"
)

func (s *service) Migrate() {
	s.pRepo.AutoMigrate()
}

func (s *service) GetProducts(ctx *context.ApplicationContext) (products []*ProductResponse, err error) {
	getProducts, err := s.pRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	products = s.toProductsResponse(getProducts)
	return
}

func (s *service) GetProduct(ctx *context.ApplicationContext, id string) (product *ProductResponse, err error) {
	getProduct, err := s.pRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	product = entityToProdcutResponse(getProduct)
	return
}
