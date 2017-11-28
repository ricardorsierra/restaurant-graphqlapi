package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/printer/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// PrintersResolver is the resolver of PrintersType
var PrintersResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO PRINTERS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("PRINTER NOT FOUND")
	}

	var printers []types.Printer

	if err := models.PrinterCollection().Find(query).Limit(30).All(&printers); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(printers) == 0 {
		return nil, errMessage
	}

	return printers, nil
}
