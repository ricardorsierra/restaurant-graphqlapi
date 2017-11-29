package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/table/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateTableResolver is the resolver of the CreateTableType
var CreateTableResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newTable, ok := params.Args["table"].(map[string]interface{})
	if !ok {
		return newTable, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newTable["password"].(string))

	table := types.Table{
		ID:       bson.NewObjectId(),
		Name:     newTable["name"].(string),
		Email:    newTable["email"].(string),
		Password: pass,
	}

	if err := models.TableCollection().Insert(table); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return table, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
