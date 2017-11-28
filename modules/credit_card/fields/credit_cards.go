package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/credit_card/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/credit_card/types"

	"github.com/graphql-go/graphql"
)

// CreditCardsType is the graphql creditCards type
var CreditCardsType = &graphql.Field{
	Type: graphql.NewList(types.CreditCardType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.CreditCardsResolver,
}
