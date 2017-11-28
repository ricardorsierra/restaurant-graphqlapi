package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// BiloCodeCollection returns the biloCode collection
func BiloCodeCollection() *mgo.Collection {
	return database.Collection("bilo_codes", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
