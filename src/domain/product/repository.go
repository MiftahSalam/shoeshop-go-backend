package product

import (
	"shoeshop-backend/src/interfaces/http/context"
)

type Repository interface {
	CountByName(ctx *context.ApplicationContext, keyword string) (total int64, err error)
	Search(ctx *context.ApplicationContext, keyword string, offset, limit int) (products []*Product, err error)
	GetAll(ctx *context.ApplicationContext) (products []*Product, err error)
	GetById(ctx *context.ApplicationContext, id string) (product *Product, err error)
	GetReviewByProductAndUser(ctx *context.ApplicationContext, productId, userId string) (review *Review, err error)
	UpdateColumn(ctx *context.ApplicationContext, product *Product) (err error)
	SaveReview(ctx *context.ApplicationContext, review *Review) (err error)
	AutoMigrate()
}
