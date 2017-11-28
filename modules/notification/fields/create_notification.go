package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/notification/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/notification/types"

	"github.com/graphql-go/graphql"
)

// CreateNotificationType is the graphql type to create post
var CreateNotificationType = &graphql.Field{
	Type: types.CreatedNotification,
	Args: graphql.FieldConfigArgument{
		"notification": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.NotificationInputType),
		},
	},
	Resolve: resolvers.CreateNotificationResolver,
}
