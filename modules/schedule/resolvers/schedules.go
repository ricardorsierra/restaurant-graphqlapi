package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/schedule/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// SchedulesResolver is the resolver of SchedulesType
var SchedulesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO SCHEDULES YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("SCHEDULE NOT FOUND")
	}

	var schedules []types.Schedule

	if err := models.ScheduleCollection().Find(query).Limit(30).All(&schedules); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(schedules) == 0 {
		return nil, errMessage
	}

	return schedules, nil
}
