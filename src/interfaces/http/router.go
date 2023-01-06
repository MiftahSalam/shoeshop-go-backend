package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"shoeshop-backend/src/interfaces/http/interceptor"
)

func setupRouter(e *echo.Echo, handler *handler, interceptor *interceptor.Interceptor) {
	e.GET("", func(e echo.Context) error {
		return e.JSON(http.StatusNotFound, "")
	})

	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Shoeshop Service... "+time.Now().Format(time.RFC3339))
	})

	groupProduct := e.Group("/api/v1/product")
	{
		groupProduct.GET("/all", handler.pHandler.GetAll, interceptor.Interceptor1)
	}

	groupAccount := e.Group("/api/v1/user")
	{
		groupAccount.GET("/profile", nil, interceptor.Interceptor2)
	}

}
