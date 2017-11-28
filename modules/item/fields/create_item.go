package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/item/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/item/types"

	"github.com/graphql-go/graphql"
)

// CreateItemType is the graphql type to create post
var CreateItemType = &graphql.Field{
	Type: types.CreatedItem,
	Args: graphql.FieldConfigArgument{
		"item": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.ItemInputType),
		},
	},
	Resolve: resolvers.CreateItemResolver,
}
