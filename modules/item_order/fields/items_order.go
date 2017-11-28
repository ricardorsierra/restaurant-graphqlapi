package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/item_order/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/item_order/types"

	"github.com/graphql-go/graphql"
)

// ItemsOrderType is the graphql itemsOrder type
var ItemsOrderType = &graphql.Field{
	Type: graphql.NewList(types.ItemOrderType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.ItemsOrderResolver,
}
