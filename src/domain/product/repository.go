package product

import (
	"shoeshop-backend/src/interfaces/http/context"
)

type Repository interface {
	GetAll(ctx *context.ApplicationContext) (products []*Product, err error)
	GetById(ctx *context.ApplicationContext, id string) (product *Product, err error)
	GetReviewByProductAndUser(ctx *context.ApplicationContext, productId, userId string) (review *Review, err error)
	SaveReview(ctx *context.ApplicationContext, review *Review) (err error)
	AutoMigrate()
}
