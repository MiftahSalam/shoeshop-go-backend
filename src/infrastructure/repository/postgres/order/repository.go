package order

import (
	"fmt"

	"shoeshop-backend/src/domain/order"
	"shoeshop-backend/src/infrastructure/database"
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
