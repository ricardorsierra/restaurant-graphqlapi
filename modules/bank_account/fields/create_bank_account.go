package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/bank_account/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/bank_account/types"

	"github.com/graphql-go/graphql"
)

// CreateBankAccountType is the graphql type to create post
var CreateBankAccountType = &graphql.Field{
	Type: types.CreatedBankAccount,
	Args: graphql.FieldConfigArgument{
		"bankAccount": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.BankAccountInputType),
		},
	},
	Resolve: resolvers.CreateBankAccountResolver,
}
