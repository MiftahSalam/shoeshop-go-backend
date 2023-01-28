package product

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/usecase/product"
)

type (
	Service interface {
		GetAllTest(ctx *context.ApplicationContext, request *CreateProductRequest) ([]*ProductResponse, error)
		GetAll(ctx *context.ApplicationContext) ([]*Product, error)
	}

	service struct {
		pUC product.Service
	}
)

func NewService(pUC product.Service) Service {
	if pUC == nil {
		panic("product usecase is nil")
	}
	return &service{pUC: pUC}
}

func (s *service) GetAllTest(ctx *context.ApplicationContext, request *CreateProductRequest) (out []*ProductResponse, err error) {
	res, err := s.pUC.GetProducts(ctx)
	if err != nil {
		return
	}
	out = ToProductsResponseTest(res)
	return
}

func (s *service) GetAll(ctx *context.ApplicationContext) (out []*Product, err error) {
	res, err := s.pUC.GetProducts(ctx)
	if err != nil {
		return
	}
	out = ToProductsResponse(res)
	return
}
