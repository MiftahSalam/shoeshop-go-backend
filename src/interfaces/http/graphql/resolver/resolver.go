package resolver

import (
	graph "shoeshop-backend/src/interfaces/http/graphql"
	productView "shoeshop-backend/src/interfaces/http/view/product"
	userView "shoeshop-backend/src/interfaces/http/view/user"
	"shoeshop-backend/src/usecase/token"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productView  productView.Service
	userView     userView.Service
	serviceToken token.Service
}

func NewResolver(vProduct productView.Service, vUser userView.Service, sToken token.Service) graph.Config {
	r := Resolver{productView: vProduct, userView: vUser, serviceToken: sToken}

	return graph.Config{Resolvers: &r}
}
