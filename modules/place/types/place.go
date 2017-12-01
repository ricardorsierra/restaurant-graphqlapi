package types

import (
	"github.com/ricardorsierra/bilo-api/modules/post/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// PlaceType is the graphql place type
var PlaceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Place",
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

// CreatedPlace is the graphql place type with token and without posts
var CreatedPlace = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreatedPlace",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"place_name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"cep": &graphql.Field{
			Type: graphql.String,
		},
		"number": &graphql.Field{
			Type: graphql.String,
		},
		"telephone": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// PlaceInputType is the graphql input post type
var PlaceInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PlaceInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"place_name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"cep": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"number": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"telephone": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

// Place is the place type
type Place struct {
	ID       	bson.ObjectId `json:"_id" bson:"_id"`
	PlaceName   string   	  `json:"place_name" bson:"place_name"`
	Description string        `json:"description" bson:"description"`
	Email    	string        `json:"email" bson:"email"`
	Password 	string        `json:"password" bson:"password"`
	UrlTickets  string 		  `json:"url_tickets" bson:"url_tickets"`
	Cep  		string 		  `json:"cep" bson:"cep"`
	Number  	string 		  `json:"number" bson:"number"`
	Telephone   string 		  `json:"telephone" bson:"telephone"`
	Token    	string        `json:"token" bson:"-"`
}
