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

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	task := &model.Task{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   input.Text,
		UserID: input.UserID,
	}
	r.tasks = append(r.tasks, task)
	return task, nil
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

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	return r.tasks, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	return r.user, nil
}

func (r *taskResolver) User(ctx context.Context, obj *model.Task) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user" + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
