package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/profile/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateProfileResolver is the resolver of the CreateProfileType
var CreateProfileResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newProfile, ok := params.Args["profile"].(map[string]interface{})
	if !ok {
		return newProfile, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newProfile["password"].(string))

	profile := types.Profile{
		ID:       bson.NewObjectId(),
		Name:     newProfile["name"].(string),
		Email:    newProfile["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(profile)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	profile.Token = tokenString

	if err := models.ProfileCollection().Insert(profile); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return profile, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
