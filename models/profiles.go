package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// ProfileCollection returns the profile collection
func ProfileCollection() *mgo.Collection {
	return database.Collection("profiles", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
