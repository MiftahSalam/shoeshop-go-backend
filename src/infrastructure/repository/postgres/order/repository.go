package order

import (
	"errors"
	"fmt"

	"shoeshop-backend/src/domain/order"
	"shoeshop-backend/src/infrastructure/database"
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/constant"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repo struct {
	master database.ORM
	slave  database.ORM
}

func NewRepository(master database.ORM, slave database.ORM) order.Repository {
	if master == nil {
		panic("please provide sql DB")
	}
	if slave == nil {
		panic("please provide sql DB Slave")
	}
	return &repo{master: master, slave: slave}
}

func (r *repo) UpdateColumn(ctx *context.ApplicationContext, data *order.Order) (err error) {
	err = r.master.Model(&order.Order{}).Where("id = ?", data.ID.String()).UpdateColumns(data)
	if err != nil {
		ctx.Logger.Error("failed order.update: " + err.Error())
		return constant.ErrorInternalServer
	}

	return nil
}

func (r *repo) Save(ctx *context.ApplicationContext, order *order.Order) (err error) {
	err = r.master.Create(order)
	if err != nil {
		ctx.Logger.Error("failed order.Save: " + err.Error())
		return constant.ErrorInternalServer
	}

	return nil
}

func (r *repo) GetById(ctx *context.ApplicationContext, id string) (order *order.Order, err error) {
	err = r.slave.Preload("Items.Product").Preload(clause.Associations).Where("id = ?", id).First(&order)
	if err == nil {
		return
	}

	ctx.Logger.Error("failed order.GetById:" + err.Error())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return order, constant.ErrorDataNotFound
	}

	return order, constant.ErrorInternalServer
}

func (r *repo) AutoMigrate() {
	err := r.master.Migrate(&order.Item{})
	if err != nil {
		fmt.Println("error auto migrate item domain with error:", err)
	}

	err = r.master.Migrate(&order.Order{})
	if err != nil {
		fmt.Println("error auto migrate order domain with error:", err)
	}
}
