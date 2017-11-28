package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/profile/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/profile/types"

	"github.com/graphql-go/graphql"
)

// CreateProfileType is the graphql type to create post
var CreateProfileType = &graphql.Field{
	Type: types.CreatedProfile,
	Args: graphql.FieldConfigArgument{
		"profile": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.ProfileInputType),
		},
	},
	Resolve: resolvers.CreateProfileResolver,
}
