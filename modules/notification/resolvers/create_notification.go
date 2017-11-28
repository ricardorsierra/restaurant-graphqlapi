package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/notification/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateNotificationResolver is the resolver of the CreateNotificationType
var CreateNotificationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newNotification, ok := params.Args["notification"].(map[string]interface{})
	if !ok {
		return newNotification, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newNotification["password"].(string))

	notification := types.Notification{
		ID:       bson.NewObjectId(),
		Name:     newNotification["name"].(string),
		Email:    newNotification["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(notification)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	notification.Token = tokenString

	if err := models.NotificationCollection().Insert(notification); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return notification, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
