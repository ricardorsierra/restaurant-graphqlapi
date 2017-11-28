package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/schedule/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/schedule/types"

	"github.com/graphql-go/graphql"
)

// CreateScheduleType is the graphql type to create post
var CreateScheduleType = &graphql.Field{
	Type: types.CreatedSchedule,
	Args: graphql.FieldConfigArgument{
		"schedule": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.ScheduleInputType),
		},
	},
	Resolve: resolvers.CreateScheduleResolver,
}
