package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/printer/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/printer/types"

	"github.com/graphql-go/graphql"
)

// CreatePrinterType is the graphql type to create post
var CreatePrinterType = &graphql.Field{
	Type: types.CreatedPrinter,
	Args: graphql.FieldConfigArgument{
		"printer": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.PrinterInputType),
		},
	},
	Resolve: resolvers.CreatePrinterResolver,
}
