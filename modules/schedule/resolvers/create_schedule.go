package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/schedule/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateScheduleResolver is the resolver of the CreateScheduleType
var CreateScheduleResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newSchedule, ok := params.Args["schedule"].(map[string]interface{})
	if !ok {
		return newSchedule, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newSchedule["password"].(string))

	schedule := types.Schedule{
		ID:       bson.NewObjectId(),
		Name:     newSchedule["name"].(string),
		Email:    newSchedule["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(schedule)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	schedule.Token = tokenString

	if err := models.ScheduleCollection().Insert(schedule); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return schedule, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
