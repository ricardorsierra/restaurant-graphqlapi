package types

import (
	"github.com/graphql-go/graphql"

	post "github.com/ricardorsierra/bilo-api/modules/post/fields"
	user "github.com/ricardorsierra/bilo-api/modules/user/fields"

	postType "github.com/ricardorsierra/bilo-api/modules/post/types"
	userType "github.com/ricardorsierra/bilo-api/modules/user/types"
)

// ViewerType is the graphql viewer type
var ViewerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Viewer",
	Fields: graphql.Fields{
		"posts": post.PostsType,
		"users": user.UsersType,
	},
})

// Viewer is the viewer type
type Viewer struct {
	Posts []postType.Post `json:"posts"`
	Users []userType.User `json:"users"`
}
