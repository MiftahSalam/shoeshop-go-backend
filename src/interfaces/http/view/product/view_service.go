package product

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/usecase/product"
)

type (
	Service interface {
		GetAllTest(ctx *context.ApplicationContext, request *CreateProductRequest) ([]*ProductResponse, error)
		GetAll(ctx *context.ApplicationContext, filter *Search) (*Products, error)
		GetById(ctx *context.ApplicationContext, id string) (*Product, error)
		CreateReview(ctx *context.ApplicationContext, userId string, review ReviewInput) (string, error)
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

func (s *service) CreateReview(ctx *context.ApplicationContext, userId string, review ReviewInput) (res string, err error) {
	res, err = s.pUC.CreateReview(ctx, userId, product.ReviewInput(review))
	if err != nil {
		return
	}

	return
}

func (s *service) GetAllTest(ctx *context.ApplicationContext, request *CreateProductRequest) (out []*ProductResponse, err error) {
	res, err := s.pUC.GetProducts(ctx, "", 0, 0)
	if err != nil {
		return
	}
	out = ToProductsResponseTest(res)
	return
}

func (s *service) GetAll(ctx *context.ApplicationContext, filter *Search) (out *Products, err error) {
	var (
		keyword = ""
		page    = 1
		limit   = 10
	)

	if filter != nil {
		keyword, page, limit = filter.validate()
	}

	res, err := s.pUC.GetProducts(ctx, keyword, page, limit)
	if err != nil {
		return
	}
	out = ToProductsResponse(res)
	return
}

func (s *service) GetById(ctx *context.ApplicationContext, id string) (out *Product, err error) {
	res, err := s.pUC.GetProduct(ctx, id)
	if err != nil {
		return
	}
	out = toProductResponse(res)
	return
}
