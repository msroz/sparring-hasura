package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/99designs/gqlgen/graphql"
	"github.com/msroz/sparring-hasura/actions/graph/generated"
	"github.com/msroz/sparring-hasura/actions/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   input.Text,
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) SingleUpload(ctx context.Context, file *graphql.Upload) (*model.File, error) {
	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}

	return &model.File{
		ID:      1,
		Name:    file.Filename,
		Content: string(content),
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	return r.user, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user" + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
