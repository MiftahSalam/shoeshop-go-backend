package order

import "shoeshop-backend/src/interfaces/http/context"

type Repository interface {
	AutoMigrate()
	GetById(ctx *context.ApplicationContext, id string) (order *Order, err error)
	GetByUserId(ctx *context.ApplicationContext, userId string) (order []*Order, err error)
	UpdateColumn(ctx *context.ApplicationContext, order *Order) (err error)
	Save(ctx *context.ApplicationContext, order *Order) (err error)
}
