package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/notification/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// NotificationsResolver is the resolver of NotificationsType
var NotificationsResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO NOTIFICATIONS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("NOTIFICATION NOT FOUND")
	}

	var notifications []types.Notification

	if err := models.NotificationCollection().Find(query).Limit(30).All(&notifications); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(notifications) == 0 {
		return nil, errMessage
	}

	return notifications, nil
}
