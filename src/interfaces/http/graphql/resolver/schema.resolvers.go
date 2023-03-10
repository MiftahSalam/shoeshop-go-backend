package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	ctxApp "shoeshop-backend/src/interfaces/http/context"
	graph "shoeshop-backend/src/interfaces/http/graphql"
	"shoeshop-backend/src/interfaces/http/view/order"
	"shoeshop-backend/src/interfaces/http/view/product"
	"shoeshop-backend/src/interfaces/http/view/user"
	"shoeshop-backend/src/shared/constant"
)

// UserRegister is the resolver for the userRegister field.
func (r *mutationResolver) UserRegister(ctx context.Context, input user.Register) (*user.User, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	user, err := r.userView.RegisterUser(appContext, input)
	if err != nil {
		return nil, err
	}

	token, err := r.serviceToken.Generate(appContext, user.ID)
	if err != nil {
		return nil, constant.ErrorInternalServer
	}

	user.Token = token

	return user, nil
}

// UpdateUserProfile is the resolver for the updateUserProfile field.
func (r *mutationResolver) UpdateUserProfile(ctx context.Context, input user.UpdateProfile) (*user.User, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	userId, err := r.serviceToken.CheckAuth(appContext)
	if err != nil {
		return nil, err
	}

	user, err := r.userView.UpdateUser(appContext, userId, input)
	if err != nil {
		return nil, err
	}

	token, err := r.serviceToken.Generate(appContext, user.ID)
	if err != nil {
		return nil, constant.ErrorInternalServer
	}

	user.Token = token

	return user, nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input order.OrderInput) (*order.OrderResponse, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	userId, err := r.serviceToken.CheckAuth(appContext)
	if err != nil {
		return nil, err
	}

	order, err := r.orderView.CreateOrder(appContext, userId, &input)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// PayOrder is the resolver for the payOrder field.
func (r *mutationResolver) PayOrder(ctx context.Context, id string, payment order.PaymentResultInput) (*order.OrderResponse, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	_, err := r.serviceToken.CheckAuth(appContext)
	if err != nil {
		return nil, err
	}

	order, err := r.orderView.PayOrder(appContext, id, payment)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// CreateProductReview is the resolver for the createProductReview field.
func (r *mutationResolver) CreateProductReview(ctx context.Context, input product.ReviewInput) (string, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	userId, err := r.serviceToken.CheckAuth(appContext)
	if err != nil {
		return "Invalid user", err
	}

	return r.productView.CreateReview(appContext, userId, input)
}

// GetProducts is the resolver for the getProducts field.
func (r *queryResolver) GetProducts(ctx context.Context, input *product.Search) (*product.Products, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	return r.productView.GetAll(appContext, input)
}

// GetProduct is the resolver for the getProduct field.
func (r *queryResolver) GetProduct(ctx context.Context, id string) (*product.Product, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	return r.productView.GetById(appContext, id)
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, input user.Login) (*user.User, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	user, err := r.userView.LoginUser(appContext, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	token, err := r.serviceToken.Generate(appContext, user.ID)
	if err != nil {
		return nil, constant.ErrorInternalServer
	}

	user.Token = token

	return user, nil
}

// GetUserProfile is the resolver for the getUserProfile field.
func (r *queryResolver) GetUserProfile(ctx context.Context) (*user.User, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	userId, err := r.serviceToken.CheckAuth(appContext)
	if err != nil {
		return nil, err
	}

	user, err := r.userView.GetUserById(appContext, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetOrder is the resolver for the getOrder field.
func (r *queryResolver) GetOrder(ctx context.Context, id string) (*order.OrderResponse, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	_, err := r.serviceToken.CheckAuth(appContext)
	if err != nil {
		return nil, err
	}

	order, err := r.orderView.GetOrder(appContext, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetUserOrders is the resolver for the getUserOrders field.
func (r *queryResolver) GetUserOrders(ctx context.Context) ([]*order.OrderResponse, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	userId, err := r.serviceToken.CheckAuth(appContext)
	if err != nil {
		return nil, err
	}

	orders, err := r.orderView.GetOrders(appContext, userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetList(ctx context.Context) (*product.Products, error) {
	appContext := ctxApp.GetAppCtxFromContext(ctx)

	return r.productView.GetAll(appContext, &product.Search{})

}
