package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	graphql1 "shoeshop-backend/src/interfaces/http/graphql"
	"shoeshop-backend/src/interfaces/http/view/product"
)

// Products is the resolver for the products field.
func (r *productsResolver) Products(ctx context.Context, obj *product.Products) ([]*product.Product, error) {
	return obj.ProductList, nil
}

// Products returns graphql1.ProductsResolver implementation.
func (r *Resolver) Products() graphql1.ProductsResolver { return &productsResolver{r} }

type productsResolver struct{ *Resolver }
