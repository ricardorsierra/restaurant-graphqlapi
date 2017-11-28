package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/order/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// OrdersResolver is the resolver of OrdersType
var OrdersResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO ORDERS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("ORDER NOT FOUND")
	}

	var orders []types.Order

	if err := models.OrderCollection().Find(query).Limit(30).All(&orders); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(orders) == 0 {
		return nil, errMessage
	}

	return orders, nil
}
