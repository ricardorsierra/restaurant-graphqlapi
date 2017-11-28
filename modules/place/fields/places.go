package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/place/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/place/types"

	"github.com/graphql-go/graphql"
)

// PlacesType is the graphql places type
var PlacesType = &graphql.Field{
	Type: graphql.NewList(types.PlaceType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.PlacesResolver,
}
