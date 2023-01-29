package user

import (
	"errors"
	"fmt"

	"shoeshop-backend/src/domain/user"
	"shoeshop-backend/src/infrastructure/database"
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/constant"

	"gorm.io/gorm"
)

type repo struct {
	master database.ORM
	slave  database.ORM
}

func NewRepository(master database.ORM, slave database.ORM) user.Repository {
	if master == nil {
		panic("please provide sql DB")
	}
	if slave == nil {
		panic("please provide sql DB Slave")
	}
	return &repo{master: master, slave: slave}
}

func (r *repo) GetByEmail(ctx *context.ApplicationContext, email string) (user *user.User, err error) {
	err = r.slave.Where("email = ?", email).First(&user)
	if err == nil {
		return
	}

	ctx.Logger.Error("failed product.GetById:" + err.Error())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, constant.ErrorDataNotFound
	}

	return user, constant.ErrorInternalServer
}

func (r *repo) AutoMigrate() {
	err := r.master.Migrate(&user.User{})
	if err != nil {
		fmt.Println("error auto migrate user domain with error:", err)
	}
}
