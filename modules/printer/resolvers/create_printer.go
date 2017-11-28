package resolvers

import (
	"crypto/sha256"
	"fmt"
	"github.com/ricardorsierra/bilo-api/helpers/auth"
	"github.com/ricardorsierra/bilo-api/models"
	"github.com/ricardorsierra/bilo-api/modules/printer/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreatePrinterResolver is the resolver of the CreatePrinterType
var CreatePrinterResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newPrinter, ok := params.Args["printer"].(map[string]interface{})
	if !ok {
		return newPrinter, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newPrinter["password"].(string))

	printer := types.Printer{
		ID:       bson.NewObjectId(),
		Name:     newPrinter["name"].(string),
		Email:    newPrinter["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(printer)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	printer.Token = tokenString

	if err := models.PrinterCollection().Insert(printer); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return printer, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
