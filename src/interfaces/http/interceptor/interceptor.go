package interceptor

import (
	"github.com/labstack/echo/v4"
)

type Interceptor struct {
}

func NewInterceptor() *Interceptor {

	return &Interceptor{}
}

func (i *Interceptor) Interceptor1(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Debug("Interceptor 1")
		return h(c)
	}
}

func (i *Interceptor) Interceptor2(h echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		c.Logger().Debug("Interceptor 2")
		return h(c)
	}

}
