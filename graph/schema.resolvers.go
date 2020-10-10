package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	auth "github.com/Salaton/screening-test.git/auth"
	"github.com/Salaton/screening-test.git/graph/generated"
	model "github.com/Salaton/screening-test.git/graph/model"
	db "github.com/Salaton/screening-test.git/postgres"
	users "github.com/Salaton/screening-test.git/users"
)

var DB db.DBClient

// var usermethod users.UserMethodsInterface

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreatedUser) (string, error) {
	// var user model.User
	DB.CreateUser(input)

	token, err := auth.CreateNewToken(input.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginDetails) (string, error) {
	var user model.User
	user.Username = input.Username
	user.Password = input.Password
	correct := DB.Authenticate()

	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := auth.CreateNewToken(input.Username)
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

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	var customers []*model.Customer
	// r.db.Find(&customers)
	// DB.FindCustomers(&customers)
	return customers, nil

	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	err := r.db.Preload("customers").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
	// panic(fmt.Errorf("not implemented"))
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

// func (r *mutationResolver) CreateToken(ctx context.Context, input model.NewUser) (string, error) {
// 	var user model.NewUser
// 	user.Username = input.Username
// 	token, err := auth.CreateNewToken(user.Username)
// 	if err != nil {
// 		return "", nil
// 	}

// 	return token, nil
// 	// panic(fmt.Errorf("not implemented"))
// }
