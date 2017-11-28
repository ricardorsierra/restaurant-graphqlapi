package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/item_type/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/item_type/types"

	"github.com/graphql-go/graphql"
)

// CreateItemTypeType is the graphql type to create post
var CreateItemTypeType = &graphql.Field{
	Type: types.CreatedItemType,
	Args: graphql.FieldConfigArgument{
		"itemType": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.ItemTypeInputType),
		},
	},
	Resolve: resolvers.CreateItemTypeResolver,
}
