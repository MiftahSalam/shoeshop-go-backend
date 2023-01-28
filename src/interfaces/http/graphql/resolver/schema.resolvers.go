package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	ctxApp "shoeshop-backend/src/interfaces/http/context"
	graph "shoeshop-backend/src/interfaces/http/graphql"
	"shoeshop-backend/src/interfaces/http/view/product"
)

// GetList is the resolver for the getList field.
func (r *productQueryResolver) GetList(ctx context.Context) ([]*product.Product, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	return r.productView.GetAll(appContext)
}

// ProductQuery returns graph.ProductQueryResolver implementation.
func (r *Resolver) ProductQuery() graph.ProductQueryResolver { return &productQueryResolver{r} }

type productQueryResolver struct{ *Resolver }
