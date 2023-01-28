package review

import (
	"fmt"

	"shoeshop-backend/src/domain/review"
	"shoeshop-backend/src/infrastructure/database"
)

type repo struct {
	master database.ORM
	slave  database.ORM
}

func NewRepository(master database.ORM, slave database.ORM) review.Repository {
	if master == nil {
		panic("please provide sql DB")
	}
	if slave == nil {
		panic("please provide sql DB Slave")
	}
	return &repo{master: master, slave: slave}
}

func (r *repo) AutoMigrate() {
	err := r.master.Migrate(&review.Review{})
	if err != nil {
		fmt.Println("error auto migrate review domain with error:", err)
	}
}
