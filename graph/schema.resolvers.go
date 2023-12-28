package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/ZainJavedDev/gqlgen-jokes-2/controllers"
	"github.com/ZainJavedDev/gqlgen-jokes-2/graph/model"
)

// Jokes is the resolver for the jokes field.
func (r *queryResolver) Jokes(ctx context.Context) ([]*model.Joke, error) {
	jokes := controllers.GetAllJokes()
	return jokes, nil
}

// Joke is the resolver for the joke field.
func (r *queryResolver) Joke(ctx context.Context, id string) (*model.Joke, error) {
	joke := controllers.GetAJoke(id)
	return joke, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
