package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// TableCollection returns the table collection
func TableCollection() *mgo.Collection {
	return database.Collection("tables", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
