package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// PrinterCollection returns the printer collection
func PrinterCollection() *mgo.Collection {
	return database.Collection("printers", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
