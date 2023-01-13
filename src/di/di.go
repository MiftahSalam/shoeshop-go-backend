package di

import (
	"shoeshop-backend/src/infrastructure/logger"
	pRepository "shoeshop-backend/src/infrastructure/repository/postgres/product"
	"shoeshop-backend/src/interfaces/http/interceptor"
	pView "shoeshop-backend/src/interfaces/http/view/product"
	"shoeshop-backend/src/shared/config"
	"shoeshop-backend/src/shared/database"
	"shoeshop-backend/src/usecase/product"
)

type DI struct {
	Configuration  *config.Configuration
	Logger         logger.Logger
	ProductService product.Service
	Interceptor    *interceptor.Interceptor
	ProductView    pView.Service
}

func Setup() *DI {
	cfg := config.Setup()

	log := logger.NewLogger(&cfg.Logger)

	dbMaster := database.Setup(cfg.Database, &log)
	dbSlave := database.Setup(cfg.Database, &log)

	pRepo := pRepository.NewRepository(dbMaster, dbSlave)

	pService := product.NewService(pRepo)

	intercept := interceptor.NewInterceptor()

	vProduct := pView.NewService(pService)

	return &DI{
		Configuration:  cfg,
		Logger:         log,
		ProductService: pService,
		Interceptor:    intercept,
		ProductView:    vProduct,
	}
}
