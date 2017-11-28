package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/bilo_code/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateBiloCodeResolver is the resolver of the CreateBiloCodeType
var CreateBiloCodeResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newBiloCode, ok := params.Args["biloCode"].(map[string]interface{})
	if !ok {
		return newBiloCode, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newBiloCode["password"].(string))

	biloCode := types.BiloCode{
		ID:       bson.NewObjectId(),
		Name:     newBiloCode["name"].(string),
		Email:    newBiloCode["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(biloCode)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	biloCode.Token = tokenString

	if err := models.BiloCodeCollection().Insert(biloCode); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return biloCode, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
