package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/bank/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateBankResolver is the resolver of the CreateBankType
var CreateBankResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newBank, ok := params.Args["bank"].(map[string]interface{})
	if !ok {
		return newBank, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newBank["password"].(string))

	bank := types.Bank{
		ID:       bson.NewObjectId(),
		Name:     newBank["name"].(string),
		Email:    newBank["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(bank)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	bank.Token = tokenString

	if err := models.BankCollection().Insert(bank); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return bank, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
