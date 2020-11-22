package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Salaton/screening-test/auth"
	"github.com/Salaton/screening-test/graph/generated"
	"github.com/Salaton/screening-test/graph/model"
	db "github.com/Salaton/screening-test/postgres"
	"github.com/Salaton/screening-test/users"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginDetails) (string, error) {
	var customer model.Customer
	customer.Email = input.Email
	customer.Password = input.Password
	correct := DB.Authenticate(input)

	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := auth.CreateNewToken(input.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	DB.CreateCustomer(input)
	return &model.Customer{}, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Order{}, fmt.Errorf("access denied, you need to log in")
	}
	DB.CreateOrder(input)
	return &model.Order{}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.ProductInput) (*model.Product, error) {
	DB.CreateProduct(input)
	return &model.Product{}, nil
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID *int, input model.OrderInput) (*model.Order, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Order{}, fmt.Errorf("access denied, you need to log in")
	}
	DB.UpdateOrder(*orderID, input)
	return &model.Order{}, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (string, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return "", fmt.Errorf("access denied, you need to log in")
	}
	DB.DeleteOrder(orderID)
	return "Order has been successfully deleted", nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	return DB.FindCustomers(), nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	return DB.FindOrders(), nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return DB.FindProducts(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var DB db.DBClient
