package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/item_order/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/item_order/types"

	"github.com/graphql-go/graphql"
)

// CreateItemOrderType is the graphql type to create post
var CreateItemOrderType = &graphql.Field{
	Type: types.CreatedItemOrder,
	Args: graphql.FieldConfigArgument{
		"itemOrder": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.ItemOrderInputType),
		},
	},
	Resolve: resolvers.CreateItemOrderResolver,
}
