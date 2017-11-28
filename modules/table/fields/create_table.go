package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/table/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/table/types"

	"github.com/graphql-go/graphql"
)

// CreateTableType is the graphql type to create post
var CreateTableType = &graphql.Field{
	Type: types.CreatedTable,
	Args: graphql.FieldConfigArgument{
		"table": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.TableInputType),
		},
	},
	Resolve: resolvers.CreateTableResolver,
}
