package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/notification/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/notification/types"

	"github.com/graphql-go/graphql"
)

// NotificationsType is the graphql notifications type
var NotificationsType = &graphql.Field{
	Type: graphql.NewList(types.NotificationType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.NotificationsResolver,
}
