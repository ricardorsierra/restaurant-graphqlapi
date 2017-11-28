package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/item_type/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/item_type/types"

	"github.com/graphql-go/graphql"
)

// ItemTypesType is the graphql itemTypes type
var ItemTypesType = &graphql.Field{
	Type: graphql.NewList(types.ItemTypeType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.ItemTypesResolver,
}
