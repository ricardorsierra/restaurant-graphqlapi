package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// NotificationCollection returns the notification collection
func NotificationCollection() *mgo.Collection {
	return database.Collection("notifications", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
