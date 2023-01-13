package http

import (
	"github.com/labstack/echo/v4"

	"shoeshop-backend/src/interfaces/http/context"
	"shoeshop-backend/src/interfaces/http/view/product"
)

type productHandler struct {
	pView product.Service
}

func SetupProductHandler(pVService product.Service) *productHandler {
	if pVService == nil {
		panic("Product view is nil")
	}
	return &productHandler{
		pView: pVService,
	}
}

func (h *productHandler) GetAll(e echo.Context) error {
	appContext := context.ParseApplicationContext(e)

	res, err := h.pView.GetAll(appContext, &product.CreateProductRequest{})
	if err != nil {
		return appContext.FailWithData(err, res)
	}

	return appContext.Success(res)
}
