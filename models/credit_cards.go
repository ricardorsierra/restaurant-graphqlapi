package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// CreditCardCollection returns the creditCard collection
func CreditCardCollection() *mgo.Collection {
	return database.Collection("credit_cards", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
