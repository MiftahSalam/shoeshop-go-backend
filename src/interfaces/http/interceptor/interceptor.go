package interceptor

import (
	"shoeshop-backend/src/interfaces/http/context"

	"github.com/labstack/echo/v4"
)

type Interceptor struct {
}

func NewInterceptor() *Interceptor {

	return &Interceptor{}
}

func (i *Interceptor) Interceptor1(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		appContext := context.ParseApplicationContext(c)
		appContext.Logger.Info("Interceptor 1")

		return h(c)
	}
}

func (i *Interceptor) Interceptor2(h echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		appContext := context.ParseApplicationContext(c)
		appContext.Logger.Info("Interceptor 2")
		return h(c)
	}

}
