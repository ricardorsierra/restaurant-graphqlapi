package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/bilo_code/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/bilo_code/types"

	"github.com/graphql-go/graphql"
)

// BiloCodesType is the graphql biloCodes type
var BiloCodesType = &graphql.Field{
	Type: graphql.NewList(types.BiloCodeType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.BiloCodesResolver,
}
