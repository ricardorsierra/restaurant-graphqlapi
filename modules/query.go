package modules

import (
	viewer "github.com/ricardorsierra/bilo-api/modules/viewer/fields"

	"github.com/graphql-go/graphql"
)

// QueryType is the root query
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"viewer": viewer.ViewerType,
	},
})
