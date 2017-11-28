package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/item_type/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// ItemTypesResolver is the resolver of ItemTypesType
var ItemTypesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO ItemTypeS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("ItemType NOT FOUND")
	}

	var itemTypes []types.ItemType

	if err := models.ItemTypeCollection().Find(query).Limit(30).All(&itemTypes); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(itemTypes) == 0 {
		return nil, errMessage
	}

	return itemTypes, nil
}
