package http

import (
	"shoeshop-backend/src/di"
)

type handler struct {
	pHandler *productHandler
}

func setupHandler(di *di.DI) *handler {
	pHandler := SetupProductHandler(di.ProductView)

	return &handler{
		pHandler: pHandler,
	}
}
