package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/profile/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// ProfilesResolver is the resolver of ProfilesType
var ProfilesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO PROFILES YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("PROFILE NOT FOUND")
	}

	var profiles []types.Profile

	if err := models.ProfileCollection().Find(query).Limit(30).All(&profiles); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(profiles) == 0 {
		return nil, errMessage
	}

	return profiles, nil
}
