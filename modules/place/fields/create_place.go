package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/place/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/place/types"

	"github.com/graphql-go/graphql"
)

// CreatePlaceType is the graphql type to create post
var CreatePlaceType = &graphql.Field{
	Type: types.CreatedPlace,
	Args: graphql.FieldConfigArgument{
		"place": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.PlaceInputType),
		},
	},
	Resolve: resolvers.CreatePlaceResolver,
}
