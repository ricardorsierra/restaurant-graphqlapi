package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// UserCollection returns the user collection
func UserCollection() *mgo.Collection {
	return database.Collection("users", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
