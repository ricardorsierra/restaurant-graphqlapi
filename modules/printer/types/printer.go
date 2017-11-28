package types

import (
	"github.com/ricardorsierra/bilo-api/modules/post/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// PrinterType is the graphql printer type
var PrinterType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Printer",
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

// CreatedPrinter is the graphql printer type with token and without posts
var CreatedPrinter = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreatedPrinter",
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

// PrinterInputType is the graphql input post type
var PrinterInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PrinterInput",
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

// Printer is the printer type
type Printer struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"-" bson:"password"`
	Posts    types.Post    `json:"posts" bson:"-"`
	Token    string        `json:"token" bson:"-"`
}
