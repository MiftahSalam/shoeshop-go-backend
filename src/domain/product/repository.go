package product

import (
	"shoeshop-backend/src/interfaces/http/context"
)

type Repository interface {
	GetAll(ctx *context.ApplicationContext) (products []*Product, err error)
	GetById(ctx *context.ApplicationContext, id string) (product *Product, err error)
	AutoMigrate()
}
