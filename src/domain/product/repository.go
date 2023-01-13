package product

import (
	"shoeshop-backend/src/interfaces/http/context"
)

type Repository interface {
	GetAll(ctx *context.ApplicationContext) (products []*Product, err error)
}
