package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"

	"shoeshop-backend/src/di"
	"shoeshop-backend/src/interfaces/http/middleware"
	"shoeshop-backend/src/interfaces/http/rest"
)

func Start(di *di.DI) {
	server := echo.New()
	gqlServer := echo.New()

	middleware.SetupMiddleware(server, di)
	middleware.SetupMiddleware(gqlServer, di)

	rest.SetupRouter(server, rest.SetupHandler(di), di.Interceptor)
	setupGQL(gqlServer, di.Interceptor, di)

	errs := errgroup.Group{}

	errs.Go(func() error {
		if err := server.Start(di.Configuration.HttpPort()); err != nil {
			panic(err)
		}
		return nil
	})

	errs.Go(func() error {
		if err := gqlServer.Start(di.Configuration.GQLHttpPort()); err != nil {
			panic(err)
		}
		return nil
	})

	if err := errs.Wait(); err != nil {
		fmt.Println(err)
	}
}
