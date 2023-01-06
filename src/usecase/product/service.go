package product

import "github.com/labstack/echo/v4"

type ProductService interface {
	GetProduct(ctx *echo.Context) (data interface{}, err error)
}
