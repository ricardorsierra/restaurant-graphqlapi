package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// ItemCollection returns the item collection
func ItemCollection() *mgo.Collection {
	return database.Collection("items", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
