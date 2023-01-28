package resolver

import (
	graph "shoeshop-backend/src/interfaces/http/graphql"
	productView "shoeshop-backend/src/interfaces/http/view/product"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productView productView.Service
}

func NewResolver(vProduct productView.Service) graph.Config {
	r := Resolver{productView: vProduct}

	return graph.Config{Resolvers: &r}
}
