package resolvers

import (
	"github.com/ricardorsierra/bilo-api/modules/viewer/types"

	"github.com/graphql-go/graphql"
)

// ViewerResolver is the resolver of ViewerType
var ViewerResolver = func(p graphql.ResolveParams) (interface{}, error) {
	return types.Viewer{}, nil
}
