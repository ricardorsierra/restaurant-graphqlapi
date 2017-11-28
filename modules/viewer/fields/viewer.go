package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/viewer/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/viewer/types"

	"github.com/graphql-go/graphql"
)

// ViewerType is the graphql viewer type
var ViewerType = &graphql.Field{
	Type: types.ViewerType,
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.ViewerResolver,
}
