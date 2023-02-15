package product

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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

func (r *repo) AutoMigrate() {
	err := r.master.Migrate(&product.Product{})
	if err != nil {
		fmt.Println("error auto migrate product domain with error:", err)
	}

	err = r.master.Migrate(&product.Review{})
	if err != nil {
		fmt.Println("error auto migrate review domain with error:", err)
	}

}

func (r *repo) GetReviewByProductAndUser(ctx *context.ApplicationContext, productId, userId string) (review *product.Review, err error) {
	err = r.slave.Where("user_id = ? and product_id = ?", userId, productId).First(&review)
	if err == nil {
		return
	}

	ctx.Logger.Error("failed review.GetByProductAndUser:" + err.Error())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return review, constant.ErrorDataNotFound
	}

	return review, constant.ErrorInternalServer
}

func (r *repo) SaveReview(ctx *context.ApplicationContext, review *product.Review) (err error) {
	err = r.master.Create(review)
	if err != nil {
		ctx.Logger.Error("failed review.Save: " + err.Error())
		return constant.ErrorInternalServer
	}

	return nil
}

func (r *repo) UpdateColumn(ctx *context.ApplicationContext, pData *product.Product) (err error) {
	err = r.master.Model(&product.Product{}).Where("id = ?", pData.ID.String()).UpdateColumns(pData)
	if err != nil {
		ctx.Logger.Error("failed order.update: " + err.Error())
		return constant.ErrorInternalServer
	}

	return nil
}

func (r *repo) CountByName(ctx *context.ApplicationContext, keyword string) (total int64, err error) {
	qFilter, qArg := searchFilter(keyword)
	err = r.slave.Where(qFilter.String(), qArg...).Model(&product.Product{}).Count(&total)
	if err != nil {
		ctx.Logger.Error("failed product.CountByName: " + err.Error())
		return 0, constant.ErrorInternalServer
	}
	return
}

func (r *repo) Search(ctx *context.ApplicationContext, keyword string, offset, limit int) (products []*product.Product, err error) {
	qFilter, qArg := searchFilter(keyword)
	query := r.slave.Preload("Reviews.User").Preload(clause.Associations).Where(qFilter.String(), qArg...)

	err = query.Offset(int64(limit * (offset - 1))).Limit(int64(limit)).Find(&products)
	if err == nil {
		return
	}

	ctx.Logger.Error("failed product.Search:" + err.Error())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return products, constant.ErrorDataNotFound
	}

	return products, constant.ErrorInternalServer
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

func (r *repo) GetById(ctx *context.ApplicationContext, id string) (product *product.Product, err error) {
	err = r.slave.Preload("Reviews.User").Preload(clause.Associations).Where("id = ?", id).First(&product)
	if err == nil {
		return
	}

	ctx.Logger.Error("failed product.GetById:" + err.Error())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return product, constant.ErrorDataNotFound
	}

	return product, constant.ErrorInternalServer
}
