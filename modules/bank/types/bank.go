package types

import (
	"github.com/ricardorsierra/bilo-api/modules/post/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// BankType is the graphql bank type
var BankType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Bank",
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

// CreatedBank is the graphql bank type with token and without posts
var CreatedBank = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreatedBank",
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

// BankInputType is the graphql input post type
var BankInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BankInput",
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

// Bank is the bank type
type Bank struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"-" bson:"password"`
	Posts    types.Post    `json:"posts" bson:"-"`
	Token    string        `json:"token" bson:"-"`
}
