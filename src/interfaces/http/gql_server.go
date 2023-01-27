package http

import (
	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"

	graph "shoeshop-backend/src/interfaces/http/graphql"
	resolver "shoeshop-backend/src/interfaces/http/graphql/resolver"
	"shoeshop-backend/src/interfaces/http/interceptor"
)

func setupGQL(e *echo.Echo, interceptor *interceptor.Interceptor) {
	server := gqlHandler.NewDefaultServer(graph.NewExecutableSchema(resolver.NewResolver()))

	e.GET("/", func(c echo.Context) error {
		playground.Handler("Shoeshop", "/graphql").ServeHTTP(c.Response(), c.Request())
		return nil
	}, interceptor.Interceptor1)
	e.POST("/graphql", func(c echo.Context) error {
		server.ServeHTTP(c.Response(), c.Request())
		return nil
	}, interceptor.Interceptor2)
}
