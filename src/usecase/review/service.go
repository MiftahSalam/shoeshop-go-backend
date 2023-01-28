package review

import (
	"shoeshop-backend/src/domain/review"
)

type (
	Service interface {
		Migrate()
	}

	service struct {
		rRepo review.Repository
	}
)

func NewService(rRepo review.Repository) Service {
	if rRepo == nil {
		panic("please provide product repository")
	}

	return &service{
		rRepo: rRepo,
	}
}
