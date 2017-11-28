package fields

import (
	"github.com/ricardorsierra/bilo-api/modules/schedule/resolvers"
	"github.com/ricardorsierra/bilo-api/modules/schedule/types"

	"github.com/graphql-go/graphql"
)

// SchedulesType is the graphql schedules type
var SchedulesType = &graphql.Field{
	Type: graphql.NewList(types.ScheduleType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.SchedulesResolver,
}
