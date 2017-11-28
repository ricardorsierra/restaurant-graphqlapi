package models

import (
	"github.com/ricardorsierra/bilo-api/helpers/database"

	mgo "gopkg.in/mgo.v2"
)

// ScheduleCollection returns the schedule collection
func ScheduleCollection() *mgo.Collection {
	return database.Collection("schedules", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
