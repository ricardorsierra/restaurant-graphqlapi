package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/order/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/order/types"

	"github.com/graphql-go/graphql"
)

// CreateOrderType is the graphql type to create post
var CreateOrderType = &graphql.Field{
	Type: types.CreatedOrder,
	Args: graphql.FieldConfigArgument{
		"order": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.OrderInputType),
		},
	},
	Resolve: resolvers.CreateOrderResolver,
}
