package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/credit_card/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreditCardsResolver is the resolver of CreditCardsType
var CreditCardsResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO CreditCardS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("CreditCard NOT FOUND")
	}

	var creditCards []types.CreditCard

	if err := models.CreditCardCollection().Find(query).Limit(30).All(&creditCards); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(creditCards) == 0 {
		return nil, errMessage
	}

	return creditCards, nil
}
