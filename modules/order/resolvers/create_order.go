package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/order/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateOrderResolver is the resolver of the CreateOrderType
var CreateOrderResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newOrder, ok := params.Args["order"].(map[string]interface{})
	if !ok {
		return newOrder, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newOrder["password"].(string))

	order := types.Order{
		ID:       bson.NewObjectId(),
		Name:     newOrder["name"].(string),
		Email:    newOrder["email"].(string),
		Password: pass,
	}

	if err := models.OrderCollection().Insert(order); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return order, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
