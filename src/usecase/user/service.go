package user

import (
	"shoeshop-backend/src/domain/user"
)

type (
	Service interface {
		Migrate()
	}

	service struct {
		uRepo user.Repository
	}
)

func NewService(uRepo user.Repository) Service {
	if uRepo == nil {
		panic("please provide product repository")
	}

	return &service{
		uRepo: uRepo,
	}
}
