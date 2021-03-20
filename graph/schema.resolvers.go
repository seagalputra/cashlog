package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/seagalputra/cashlog/graph/generated"
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/seagalputra/cashlog/internal/auth"
	"github.com/seagalputra/cashlog/internal/transaction"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.AuthPayload, error) {
	login := model.Login{
		Username: username,
		Password: password,
	}

	payload, err := r.UserService.Authenticate(login)
	if err != nil {
		return &model.AuthPayload{}, err
	}

	return payload, nil
}

func (r *mutationResolver) Register(ctx context.Context, newUser model.RegisterUser) (*model.AuthPayload, error) {
	payload, err := r.UserService.RegisterAccount(newUser)
	if err != nil {
		return &model.AuthPayload{}, err
	}

	return payload, nil
}

func (r *mutationResolver) CreateTransaction(ctx context.Context, newTransaction model.CreateTransaction) (*model.Transaction, error) {
	userID := auth.ForContext(ctx)
	if userID == "" {
		return &model.Transaction{}, fmt.Errorf("access denied")
	}

	var err error
	req := transaction.TransactionReq{
		UserID: userID,
		Trx:    newTransaction,
	}
	trx, err := r.TransactionService.CreateTransaction(req)
	if err != nil {
		return &model.Transaction{}, err
	}

	return trx, nil
}

func (r *queryResolver) Transaction(ctx context.Context, transactionID string) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Transactions(ctx context.Context, from string, to string) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
