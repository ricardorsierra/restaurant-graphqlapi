package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/bilo_code/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/bilo_code/types"

	"github.com/graphql-go/graphql"
)

// CreateBiloCodeType is the graphql type to create post
var CreateBiloCodeType = &graphql.Field{
	Type: types.CreatedBiloCode,
	Args: graphql.FieldConfigArgument{
		"biloCode": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.BiloCodeInputType),
		},
	},
	Resolve: resolvers.CreateBiloCodeResolver,
}
