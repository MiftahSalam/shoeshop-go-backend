package di

import (
	"shoeshop-backend/src/infrastructure/logger"
	oRepository "shoeshop-backend/src/infrastructure/repository/postgres/order"
	pRepository "shoeshop-backend/src/infrastructure/repository/postgres/product"
	rRepository "shoeshop-backend/src/infrastructure/repository/postgres/review"
	uRepository "shoeshop-backend/src/infrastructure/repository/postgres/user"
	"shoeshop-backend/src/interfaces/http/interceptor"
	pView "shoeshop-backend/src/interfaces/http/view/product"
	uView "shoeshop-backend/src/interfaces/http/view/user"
	"shoeshop-backend/src/shared/config"
	"shoeshop-backend/src/shared/database"
	"shoeshop-backend/src/usecase/order"
	"shoeshop-backend/src/usecase/product"
	"shoeshop-backend/src/usecase/review"
	"shoeshop-backend/src/usecase/token"
	"shoeshop-backend/src/usecase/user"
)

type DI struct {
	Configuration  *config.Configuration
	Logger         logger.Logger
	ProductService product.Service
	TokenService   token.Service
	Interceptor    *interceptor.Interceptor
	ProductView    pView.Service
	UserView       uView.Service
}

func Setup() *DI {
	cfg := config.Setup()

	log := logger.NewLogger(&cfg.Logger)

	dbMaster := database.Setup(cfg.Database, &log)
	dbSlave := database.Setup(cfg.Database, &log)

	oRepo := oRepository.NewRepository(dbMaster, dbSlave)
	pRepo := pRepository.NewRepository(dbMaster, dbSlave)
	rRepo := rRepository.NewRepository(dbMaster, dbSlave)
	uRepo := uRepository.NewRepository(dbMaster, dbSlave)

	oService := order.NewService(oRepo)
	pService := product.NewService(pRepo)
	rService := review.NewService(rRepo)
	tService := token.NewService(&cfg.Application)
	uService := user.NewService(uRepo)

	if cfg.Database.AutoMigrate {
		pService.Migrate()
		uService.Migrate()
		rService.Migrate()
		oService.Migrate()
	}

	intercept := interceptor.NewInterceptor()

	vProduct := pView.NewService(pService)
	vUser := uView.NewService(uService)

	return &DI{
		Configuration:  cfg,
		Logger:         log,
		ProductService: pService,
		TokenService:   tService,
		Interceptor:    intercept,
		ProductView:    vProduct,
		UserView:       vUser,
	}
}
