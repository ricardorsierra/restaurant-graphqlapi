package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/bank/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// BanksResolver is the resolver of BanksType
var BanksResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO BANKS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("BANK NOT FOUND")
	}

	var banks []types.Bank

	if err := models.BankCollection().Find(query).Limit(30).All(&banks); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(banks) == 0 {
		return nil, errMessage
	}

	return banks, nil
}
