package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/item_order/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateItemOrderResolver is the resolver of the CreateItemOrderType
var CreateItemOrderResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newItemOrder, ok := params.Args["itemOrder"].(map[string]interface{})
	if !ok {
		return newItemOrder, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newItemOrder["password"].(string))

	itemOrder := types.ItemOrder{
		ID:       bson.NewObjectId(),
		Name:     newItemOrder["name"].(string),
		Email:    newItemOrder["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(itemOrder)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	itemOrder.Token = tokenString

	if err := models.ItemOrderCollection().Insert(itemOrder); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return itemOrder, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
