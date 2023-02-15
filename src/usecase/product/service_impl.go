package product

import (
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/constant"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func (s *service) Migrate() {
	s.pRepo.AutoMigrate()
}

func (s *service) CreateReview(ctx *context.ApplicationContext, userId string, review ReviewInput) (string, error) {
	rData, err := s.pRepo.GetReviewByProductAndUser(ctx, review.ProductID, userId)
	if err == nil && rData.ID != uuid.Nil {
		return "Product already reviewd", constant.ErrorDataAlreadyExist
	} else if err != constant.ErrorDataNotFound {
		return "Error when trying to get review", err
	}

	user, err := s.uRepo.GetById(ctx, userId)
	if err != nil {
		return "Error when trying to get user", err
	}

	product, err := s.pRepo.GetById(ctx, review.ProductID)
	if err != nil {
		return "Error when trying to get product", err
	}

	dReview := review.ToReviewDomain(user, product)
	err = s.pRepo.SaveReview(ctx, dReview)
	if err != nil {
		return "Error when trying to create review", err
	}
	product.Reviews = append(product.Reviews, dReview)

	err = s.pRepo.UpdateColumn(ctx, s.updateRating(product))
	if err != nil {
		return "Error when trying to update product review", err
	}

	return "Review created", nil
}

func (s *service) GetProducts(ctx *context.ApplicationContext, keyword string, page, limit int) (products *ProductsResponse, err error) {
	var (
		eg          = new(errgroup.Group)
		getProducts []*product.Product
		count       int64
	)

	eg.Go(func() error {
		getProducts, err = s.pRepo.Search(ctx, keyword, page, limit)
		return err
	})

	eg.Go(func() error {
		count, err = s.pRepo.CountByName(ctx, keyword)
		return err
	})

	if err = eg.Wait(); err != nil {
		return
	}

	products = s.toProductsResponse(getProducts, count)
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

func (s *service) updateRating(pData *product.Product) *product.Product {
	var rating int
	for _, r := range pData.Reviews {
		rating += int(r.Rating)
	}
	rating /= len(pData.Reviews)

	return &product.Product{ID: pData.ID, Rating: int8(rating)}
}
