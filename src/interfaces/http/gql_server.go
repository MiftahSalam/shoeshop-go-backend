package http

import (
	"context"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"

	"shoeshop-backend/src/di"
	appContext "shoeshop-backend/src/interfaces/http/context"
	graph "shoeshop-backend/src/interfaces/http/graphql"
	resolver "shoeshop-backend/src/interfaces/http/graphql/resolver"
	"shoeshop-backend/src/interfaces/http/interceptor"
	"shoeshop-backend/src/shared/constant"
)

func setupGQL(e *echo.Echo, interceptor *interceptor.Interceptor, di *di.DI) {
	server := gqlHandler.NewDefaultServer(graph.NewExecutableSchema(resolver.NewResolver(di.ProductView, di.UserView, di.TokenService)))

	e.GET("/", func(c echo.Context) error {
		playground.Handler("Shoeshop", "/graphql").ServeHTTP(c.Response(), c.Request())
		return nil
	}, interceptor.Interceptor1)
	e.POST("/graphql", func(c echo.Context) error {
		curCtx := c.Request().Context()
		curCtx = context.WithValue(curCtx, constant.AppCtxName, appContext.ParseApplicationContext(c))
		req := c.Request().WithContext(curCtx)
		server.ServeHTTP(c.Response(), req)
		return nil
	}, interceptor.Interceptor2)
}
