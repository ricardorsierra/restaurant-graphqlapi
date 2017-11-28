package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/bank_account/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateBankAccountResolver is the resolver of the CreateBankAccountType
var CreateBankAccountResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newBankAccount, ok := params.Args["bankAccount"].(map[string]interface{})
	if !ok {
		return newBankAccount, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newBankAccount["password"].(string))

	bankAccount := types.BankAccount{
		ID:       bson.NewObjectId(),
		Name:     newBankAccount["name"].(string),
		Email:    newBankAccount["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(bankAccount)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	bankAccount.Token = tokenString

	if err := models.BankAccountCollection().Insert(bankAccount); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return bankAccount, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
