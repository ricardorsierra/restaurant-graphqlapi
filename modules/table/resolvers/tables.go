package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/table/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// TablesResolver is the resolver of TablesType
var TablesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO TABLES YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("TABLE NOT FOUND")
	}

	var tables []types.Table

	if err := models.TableCollection().Find(query).Limit(30).All(&tables); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(tables) == 0 {
		return nil, errMessage
	}

	return tables, nil
}
