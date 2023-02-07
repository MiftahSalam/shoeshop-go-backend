package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"shoeshop-backend/src/di"
	"shoeshop-backend/src/interfaces/http/context"
)

func SetupMiddleware(e *echo.Echo, di *di.DI) {
	contextInjector, _ := NewContextInjectorMiddleware(di.Logger, di.Configuration)

	e.Use(contextInjector.Injector)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "content-type", "Authorization"},
	}))
	e.HTTPErrorHandler = errorHandler
}

func errorHandler(err error, e echo.Context) {
	appContext := context.ParseApplicationContext(e)

	err = appContext.Fail(err)
	if err != nil {
		fmt.Println(err.Error())
	}
}
