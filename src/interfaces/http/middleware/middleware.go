package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"shoeshop-backend/src/di"
	"shoeshop-backend/src/interfaces/http/context"
)

func SetupMiddleware(e *echo.Echo, di *di.DI) {
	contextInjector, _ := NewContextInjectorMiddleware(di.Logger)

	e.Use(contextInjector.Injector)

	e.HTTPErrorHandler = errorHandler
}

func errorHandler(err error, e echo.Context) {
	appContext, err := context.NewApplicationContext(e)

	err = appContext.Fail(err)
	if err != nil {
		fmt.Println(err.Error())
	}
}
