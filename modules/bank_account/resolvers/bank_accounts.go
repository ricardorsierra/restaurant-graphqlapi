package resolvers

import (
	"fmt"
	"log"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/bank_account/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// BankAccountsResolver is the resolver of BankAccountsType
var BankAccountsResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO BankAccountS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("BankAccount NOT FOUND")
	}

	var bankAccounts []types.BankAccount

	if err := models.BankAccountCollection().Find(query).Limit(30).All(&bankAccounts); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(bankAccounts) == 0 {
		return nil, errMessage
	}

	return bankAccounts, nil
}
