package types

import (
	"github.com/ricardorsierra/bilo-api/modules/post/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// ItemOrderType is the graphql itemOrder type
var ItemOrderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ItemOrder",
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

// CreatedItemOrder is the graphql itemOrder type with token and without posts
var CreatedItemOrder = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreatedItemOrder",
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

// ItemOrderInputType is the graphql input post type
var ItemOrderInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ItemOrderInput",
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

// ItemOrder is the itemOrder type
type ItemOrder struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"-" bson:"password"`
	Posts    types.Post    `json:"posts" bson:"-"`
	Token    string        `json:"token" bson:"-"`
}
