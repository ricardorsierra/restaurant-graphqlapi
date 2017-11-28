package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/bank/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/bank/types"

	"github.com/graphql-go/graphql"
)

// CreateBankType is the graphql type to create post
var CreateBankType = &graphql.Field{
	Type: types.CreatedBank,
	Args: graphql.FieldConfigArgument{
		"bank": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.BankInputType),
		},
	},
	Resolve: resolvers.CreateBankResolver,
}
