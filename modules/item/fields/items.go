package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/item/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/item/types"

	"github.com/graphql-go/graphql"
)

// ItemsType is the graphql items type
var ItemsType = &graphql.Field{
	Type: graphql.NewList(types.ItemType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.ItemsResolver,
}
