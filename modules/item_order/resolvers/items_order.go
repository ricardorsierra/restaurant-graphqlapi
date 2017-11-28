package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/item_order/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// ItemsOrderResolver is the resolver of ItemsOrderType
var ItemsOrderResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO ItemsOrder YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("ItemOrder NOT FOUND")
	}

	var itemsOrder []types.ItemOrder

	if err := models.ItemOrderCollection().Find(query).Limit(30).All(&itemsOrder); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(itemsOrder) == 0 {
		return nil, errMessage
	}

	return itemsOrder, nil
}
