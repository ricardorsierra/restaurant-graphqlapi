package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/bank_account/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/bank_account/types"

	"github.com/graphql-go/graphql"
)

// BankAccountsType is the graphql bankAccounts type
var BankAccountsType = &graphql.Field{
	Type: graphql.NewList(types.BankAccountType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.BankAccountsResolver,
}
