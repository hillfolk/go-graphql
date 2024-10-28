package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"strconv"

	"github.com/hillfolk/go-graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoList []*model.Todo
	LastID   int
	UserList []*model.User
}

func (r *Resolver) id() string {
	r.LastID++
	return strconv.Itoa(r.LastID)
}

func (r *Resolver) GetUserByID(id string) *model.User {
	for _, user := range r.UserList {
		if user.ID == id {
			return user
		}
	}
	return nil
}
