package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/order/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/order/types"

	"github.com/graphql-go/graphql"
)

// OrdersType is the graphql orders type
var OrdersType = &graphql.Field{
	Type: graphql.NewList(types.OrderType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.OrdersResolver,
}
