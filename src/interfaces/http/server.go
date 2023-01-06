package http

import (
	"github.com/labstack/echo/v4"

	"shoeshop-backend/src/di"
	"shoeshop-backend/src/interfaces/http/middleware"
)

func Start(di *di.DI) {
	server := echo.New()

	middleware.SetupMiddleware(server, di)
	setupRouter(server, setupHandler(di), di.Interceptor)

	if err := server.Start(di.Configuration.HttpPort()); err != nil {
		panic(err)
	}
}
