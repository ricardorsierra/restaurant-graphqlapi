package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/place/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreatePlaceResolver is the resolver of the CreatePlaceType
var CreatePlaceResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newPlace, ok := params.Args["place"].(map[string]interface{})
	if !ok {
		return newPlace, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newPlace["password"].(string))

	place := types.Place{
		ID:       bson.NewObjectId(),
		Name:     newPlace["name"].(string),
		Email:    newPlace["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(place)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	place.Token = tokenString

	if err := models.PlaceCollection().Insert(place); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return place, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
