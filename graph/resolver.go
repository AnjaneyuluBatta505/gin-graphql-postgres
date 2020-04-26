package graph

import "github.com/AnjaneyuluBatta505/gin-graphql-postgres/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	questions []*model.Question
	choices   []*model.Choice
}
