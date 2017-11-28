package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/place/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// PlacesResolver is the resolver of PlacesType
var PlacesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO PLACES YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("PLACE NOT FOUND")
	}

	var places []types.Place

	if err := models.PlaceCollection().Find(query).Limit(30).All(&places); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(places) == 0 {
		return nil, errMessage
	}

	return places, nil
}
