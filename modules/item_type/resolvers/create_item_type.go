package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/item_type/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateItemTypeResolver is the resolver of the CreateItemTypeType
var CreateItemTypeResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newItemType, ok := params.Args["itemType"].(map[string]interface{})
	if !ok {
		return newItemType, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newItemType["password"].(string))

	itemType := types.ItemType{
		ID:       bson.NewObjectId(),
		Name:     newItemType["name"].(string),
		Email:    newItemType["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(itemType)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	itemType.Token = tokenString

	if err := models.ItemTypeCollection().Insert(itemType); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return itemType, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
