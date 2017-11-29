package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/item/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateItemResolver is the resolver of the CreateItemType
var CreateItemResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newItem, ok := params.Args["item"].(map[string]interface{})
	if !ok {
		return newItem, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newItem["password"].(string))

	item := types.Item{
		ID:       bson.NewObjectId(),
		Name:     newItem["name"].(string),
		Email:    newItem["email"].(string),
		Password: pass,
	}

	if err := models.ItemCollection().Insert(item); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return item, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
