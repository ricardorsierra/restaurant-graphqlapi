package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/credit_card/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateCreditCardResolver is the resolver of the CreateCreditCardType
var CreateCreditCardResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newCreditCard, ok := params.Args["creditCard"].(map[string]interface{})
	if !ok {
		return newCreditCard, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newCreditCard["password"].(string))

	creditCard := types.CreditCard{
		ID:       bson.NewObjectId(),
		Name:     newCreditCard["name"].(string),
		Email:    newCreditCard["email"].(string),
		Password: pass,
	}

	if err := models.CreditCardCollection().Insert(creditCard); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return creditCard, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
