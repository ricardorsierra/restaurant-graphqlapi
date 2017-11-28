package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// BankAccountCollection returns the bankAccount collection
func BankAccountCollection() *mgo.Collection {
	return database.Collection("bank_accounts", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
