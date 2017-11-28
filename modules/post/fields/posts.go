package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/post/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/post/types"

	"github.com/graphql-go/graphql"
)

// PostsType is the graphql posts type
var PostsType = &graphql.Field{
	Type: graphql.NewList(types.PostType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.PostsResolver,
}
