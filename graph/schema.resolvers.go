package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/seagalputra/cashlog/graph/generated"
	"github.com/seagalputra/cashlog/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.AuthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, newUser model.RegisterUser) (*model.AuthPayload, error) {
	auth, err := r.UserService.RegisterAccount(newUser)
	if err != nil {
		return &model.AuthPayload{}, err
	}

	return auth, nil
}

func (r *mutationResolver) CreateTransaction(ctx context.Context, newTransaction model.CreateTransaction) (*model.Transaction, error) {
	var err error

	transaction, err := r.TransactionService.CreateTransaction(newTransaction)
	if err != nil {
		return &model.Transaction{}, err
	}

	return transaction, nil
}

func (r *queryResolver) Transaction(ctx context.Context, transactionID string) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
