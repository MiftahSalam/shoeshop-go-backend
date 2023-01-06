package product

import "github.com/labstack/echo/v4"

type serviceImpl struct {
}

func Setup() *serviceImpl {
	return &serviceImpl{}
}

func (s *serviceImpl) GetProduct(ctx *echo.Context) (data interface{}, err error) {
	return
}
