package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/bilo_code/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// BiloCodesResolver is the resolver of BiloCodesType
var BiloCodesResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO BiloCodeS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("BiloCode NOT FOUND")
	}

	var biloCodes []types.BiloCode

	if err := models.BiloCodeCollection().Find(query).Limit(30).All(&biloCodes); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(biloCodes) == 0 {
		return nil, errMessage
	}

	return biloCodes, nil
}
