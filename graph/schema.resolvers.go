package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/seagalputra/cashlog/graph/generated"
	"github.com/seagalputra/cashlog/graph/model"
	"github.com/seagalputra/cashlog/internal/user"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterUser) (*model.RegisterPayload, error) {
	request := user.RegisterRequest{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Password:  input.Password,
		Email:     input.Email,
	}

	account, err := r.UserService.RegisterAccount(request)
	if err != nil {
		return &model.RegisterPayload{}, err
	}

	return &model.RegisterPayload{
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Username:  account.Username,
		Email:     account.Email,
	}, nil
}

func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.CreateTransaction) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
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
