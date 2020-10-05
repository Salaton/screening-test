package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Salaton/screening-test.git/graph/generated"
	"github.com/Salaton/screening-test.git/graph/model"
	db "github.com/Salaton/screening-test.git/postgres"
)

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	DB.CreateCustomer(input)
	return &model.Customer{}, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	DB.CreateOrder(input)
	return &model.Order{}, nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, input model.ItemInput) (*model.Item, error) {
	DB.CreateItem(input)
	return &model.Item{}, nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	var customers []*model.Customer
	r.db.Find(&customers)
	return customers, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	r.db.Preload("Item").Find(&orders)
	return orders, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item
	// similar to SELECT * FROM items
	r.db.Find(&items)
	return items, nil
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
var DB db.DBClient
