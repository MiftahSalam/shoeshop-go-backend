package product

import (
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/constant"
)

func (s *service) Migrate() {
	s.pRepo.AutoMigrate()
}

func (s *service) CreateReview(ctx *context.ApplicationContext, userId string, review ReviewInput) (string, error) {
	_, err := s.pRepo.GetReviewByProductAndUser(ctx, review.ProductID, userId)
	if err == nil {
		return "Product already reviewd", constant.ErrorDataAlreadyExist
	}

	user, err := s.uRepo.GetById(ctx, userId)
	if err != nil {
		return "Error when trying to get user", err
	}

	product, err := s.pRepo.GetById(ctx, review.ProductID)
	if err != nil {
		return "Error when trying to get product", err
	}

	err = s.pRepo.SaveReview(ctx, review.ToReviewDomain(user, product))
	if err != nil {
		return "Error when trying to create review", err
	}

	return "Review created", nil
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
