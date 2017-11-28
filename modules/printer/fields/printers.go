package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/printer/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/printer/types"

	"github.com/graphql-go/graphql"
)

// PrintersType is the graphql printers type
var PrintersType = &graphql.Field{
	Type: graphql.NewList(types.PrinterType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.PrintersResolver,
}
