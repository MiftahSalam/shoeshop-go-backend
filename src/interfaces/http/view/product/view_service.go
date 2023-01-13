package product

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/usecase/product"
)

type (
	Service interface {
		GetAll(ctx *context.ApplicationContext, request *CreateProductRequest) ([]*ProductResponse, error)
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

func (s *service) GetAll(ctx *context.ApplicationContext, request *CreateProductRequest) (out []*ProductResponse, err error) {
	res, err := s.pUC.GetProducts(ctx)
	if err != nil {
		return
	}
	out = ToProductsResponse(res)
	return
}
