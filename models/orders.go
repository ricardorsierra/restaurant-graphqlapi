package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// OrderCollection returns the order collection
func OrderCollection() *mgo.Collection {
	return database.Collection("orders", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
