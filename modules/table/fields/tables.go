package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/table/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/table/types"

	"github.com/graphql-go/graphql"
)

// TablesType is the graphql tables type
var TablesType = &graphql.Field{
	Type: graphql.NewList(types.TableType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.TablesResolver,
}
