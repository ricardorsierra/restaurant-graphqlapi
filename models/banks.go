package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// BankCollection returns the bank collection
func BankCollection() *mgo.Collection {
	return database.Collection("banks", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
