package http

import (
	"github.com/labstack/echo/v4"

	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/usecase/product"
)

type productHandler struct {
	pService product.ProductService
}

func SetupProductHandler(pService product.ProductService) *productHandler {
	if pService == nil {
		panic("Product service is nil")
	}
	return &productHandler{
		pService: pService,
	}
}

func (h *productHandler) GetAll(e echo.Context) error {
	appContext := context.ParseApplicationContext(e)

	res, err := h.pService.GetProduct(&e)
	if err != nil {
		return appContext.FailWithData(err, res)
	}

	return appContext.Success(res)
}
