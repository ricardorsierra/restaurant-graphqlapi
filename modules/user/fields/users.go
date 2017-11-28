package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/user/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/user/types"

	"github.com/graphql-go/graphql"
)

// UsersType is the graphql users type
var UsersType = &graphql.Field{
	Type: graphql.NewList(types.UserType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.UsersResolver,
}
