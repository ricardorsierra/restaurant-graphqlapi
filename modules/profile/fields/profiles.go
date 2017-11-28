package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/profile/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/profile/types"

	"github.com/graphql-go/graphql"
)

// ProfilesType is the graphql profiles type
var ProfilesType = &graphql.Field{
	Type: graphql.NewList(types.ProfileType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.ProfilesResolver,
}
