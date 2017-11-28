package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/credit_card/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/credit_card/types"

	"github.com/graphql-go/graphql"
)

// CreateCreditCardType is the graphql type to create post
var CreateCreditCardType = &graphql.Field{
	Type: types.CreatedCreditCard,
	Args: graphql.FieldConfigArgument{
		"creditCard": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.CreditCardInputType),
		},
	},
	Resolve: resolvers.CreateCreditCardResolver,
}
