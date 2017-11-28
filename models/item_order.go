package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// ItemOrderCollection returns the itemOrder collection
func ItemOrderCollection() *mgo.Collection {
	return database.Collection("item_orders", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
