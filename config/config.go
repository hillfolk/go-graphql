package config

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"

	"github.com/hillfolk/go-graphql/graph"
	"github.com/hillfolk/go-graphql/graph/model"
)

var (
	you  = &model.User{ID: "1", Name: "You"}
	them = &model.User{ID: "2", Name: "Them"}
)

type ckey string

func getUserId(ctx context.Context) string {
	return you.ID
}

func New() graph.Config {
	c := graph.Config{
		Resolvers: &graph.Resolver{
			TodoList: []*model.Todo{
				{ID: "1", Text: "A todo not to forget", Done: false, User: you},
				{ID: "2", Text: "This is the most important", Done: false, User: you},
				{ID: "3", Text: "Somebody else's todo", Done: true, User: you},
				{ID: "4", Text: "Please do this or else", Done: false, User: you},
			},
			LastID: 4,
			UserList: []*model.User{
				you,
				them,
			},
		}}

	c.Directives.HasRole = func(ctx context.Context, obj any, next graphql.Resolver, role model.Role) (any, error) {
		switch role {
		case model.RoleAdmin:
			// No admin for you!
			return nil, nil
		case model.RoleOwner:
			ownable, isOwnable := obj.(model.Ownable)
			if !isOwnable {
				return nil, errors.New("obj cant be owned")
			}

			if ownable.Owner().ID != getUserId(ctx) {
				return nil, errors.New("you don't own that")
			}
		}

		return next(ctx)
	}
	c.Directives.User = func(ctx context.Context, obj any, next graphql.Resolver, id string) (any, error) {
		return next(context.WithValue(ctx, ckey("userId"), id))
	}
	return c
}
