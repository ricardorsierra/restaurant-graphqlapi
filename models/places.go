package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// PlaceCollection returns the place collection
func PlaceCollection() *mgo.Collection {
	return database.Collection("places", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
