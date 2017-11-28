package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// ItemTypeCollection returns the itemType collection
func ItemTypeCollection() *mgo.Collection {
	return database.Collection("item_types", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
