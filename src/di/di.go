package di

import (
	"shoeshop-backend/src/infrastructure/logger"
	"shoeshop-backend/src/interfaces/http/interceptor"
	"shoeshop-backend/src/shared/config"
	"shoeshop-backend/src/usecase/product"
)

type DI struct {
	Configuration  *config.Configuration
	Logger         logger.Logger
	ProductService product.ProductService
	Interceptor    *interceptor.Interceptor
}

func Setup() *DI {
	cfg := config.Setup()

	log := logger.NewLogger(&cfg.Logger)

	pService := product.Setup()

	intercept := interceptor.NewInterceptor()

	return &DI{
		Configuration:  cfg,
		Logger:         log,
		ProductService: pService,
		Interceptor:    intercept,
	}
}
