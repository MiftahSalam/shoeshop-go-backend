package product

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/infrastructure/database"
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/constant"
)

type repo struct {
	master database.ORM
	slave  database.ORM
}

func NewRepository(master database.ORM, slave database.ORM) product.Repository {
	if master == nil {
		panic("please provide sql DB")
	}
	if slave == nil {
		panic("please provide sql DB Slave")
	}
	return &repo{master: master, slave: slave}
}

func (r *repo) GetAll(ctx *context.ApplicationContext) (products []*product.Product, err error) {
	err = r.slave.Find(&products)
	if err == nil {
		return
	}

	ctx.Logger.Error("failed product.GetAll:" + err.Error())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return products, constant.ErrorDataNotFound
	}

	return products, constant.ErrorInternalServer
}
