package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/seagalputra/cashlog/graph/generated"
	"github.com/seagalputra/cashlog/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, newUser model.RegisterUser) (*model.AuthPayload, error) {
	request := model.RegisterUser{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Username:  newUser.Username,
		Password:  newUser.Password,
		Email:     newUser.Email,
	}

	_, err := r.UserService.RegisterAccount(request)
	if err != nil {
		return &model.AuthPayload{}, err
	}

	return &model.AuthPayload{}, nil
}

func (r *mutationResolver) CreateTransaction(ctx context.Context, newTransaction model.CreateTransaction) (*model.Transaction, error) {
	var err error

	switch newTransaction.Status {
	case model.TransactionStatusIncome:
		err = r.TransactionService.CreateIncome(newTransaction)
	case model.TransactionStatusOutcome:
		err = r.TransactionService.CreateOutcome(newTransaction)
	case model.TransactionStatusWaiting:
		err = r.TransactionService.CreateWaiting(newTransaction)
	}

	if err != nil {
		return &model.Transaction{}, err
	}

	return &model.Transaction{
		Title:           newTransaction.Title,
		Amount:          newTransaction.Amount,
		TransactionDate: newTransaction.TransactionDate,
		Detail:          &model.TransactionDetail{},
		User:            &model.User{},
	}, nil
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
