package order_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"shoeshop-backend/src/infrastructure/logger"
	oRepository "shoeshop-backend/src/infrastructure/repository/postgres/order"
	pRepository "shoeshop-backend/src/infrastructure/repository/postgres/product"
	uRepository "shoeshop-backend/src/infrastructure/repository/postgres/user"
	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/shared/config"
	"shoeshop-backend/src/shared/database"
	"shoeshop-backend/src/usecase/order"
)

func TestCreateOrder(t *testing.T) {
	asserts := assert.New(t)

	os.Setenv("CONFIG_FILE", "../../../.env")
	appConfig := config.Setup()

	log := logger.NewLogger(&appConfig.Logger)

	dbMaster := database.Setup(appConfig.Database, &log)
	dbSlave := database.Setup(appConfig.Database, &log)

	oRepo := oRepository.NewRepository(dbMaster, dbSlave)
	pRepo := pRepository.NewRepository(dbMaster, dbSlave)
	uRepo := uRepository.NewRepository(dbMaster, dbSlave)

	oService := order.NewService(oRepo, pRepo, uRepo)

	e := echo.New()
	appCtx := context.NewApplicationContext(e.AcquireContext(), log)

	product1, err := pRepo.GetById(appCtx, "9b5cd22c-5220-4f42-8fcf-71c5169923b6")
	asserts.NoError(err)

	item1 := &order.Item{
		ProductId: "9b5cd22c-5220-4f42-8fcf-71c5169923b6",
		Quantity:  2,
		Price:     product1.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	product2, err := pRepo.GetById(appCtx, "edd5c297-66c9-4e27-908b-cfbdc1801186")
	asserts.NoError(err)

	item2 := &order.Item{
		ProductId: "edd5c297-66c9-4e27-908b-cfbdc1801186",
		Quantity:  12,
		Price:     product2.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	orderReq := &order.OrderRequest{
		Items: []*order.Item{item1, item2},
		ShippingAddress: order.Shipping{
			Address:    "abc",
			City:       "def",
			PostalCode: "123456",
			Country:    "ghi",
		},
		PaymentMethod: "paypal",
		TaxPrice:      1.2,
		ShippingPrice: 12.5,
		TotalPrice:    150.2,
	}

	oResp, err := oService.CreateOrder(appCtx, "023511fb-a639-4685-8360-8e74a1c4883b", orderReq)
	asserts.NoError(err)

	fmt.Println(oResp)

}
