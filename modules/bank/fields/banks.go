package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/bank/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/bank/types"

	"github.com/graphql-go/graphql"
)

// BanksType is the graphql banks type
var BanksType = &graphql.Field{
	Type: graphql.NewList(types.BankType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.BanksResolver,
}
