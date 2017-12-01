package types

import (
	"github.com/ricardorsierra/bilo-api/modules/post/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// ProfileType is the graphql profile type
var ProfileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Profile",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(types.PostType),
		},
	},
})

// CreatedProfile is the graphql profile type with token and without posts
var CreatedProfile = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreatedProfile",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// ProfileInputType is the graphql input post type
var ProfileInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ProfileInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

// Profile is the profile type
type Profile struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"-" bson:"password"`
	Posts    types.Post    `json:"posts" bson:"-"`
	Token    string        `json:"token" bson:"-"`
}
