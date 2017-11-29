package modules

import (
	post "github.com/ricardorsierra/bilo-api/modules/post/fields"
	user "github.com/ricardorsierra/bilo-api/modules/user/fields"


	biloCode "github.com/ricardorsierra/bilo-api/modules/bilo_code/fields"
	place "github.com/ricardorsierra/bilo-api/modules/place/fields"
	printer "github.com/ricardorsierra/bilo-api/modules/printer/fields"
	bankAccount "github.com/ricardorsierra/bilo-api/modules/bank_account/fields"
	creditCard "github.com/ricardorsierra/bilo-api/modules/credit_card/fields"
	item "github.com/ricardorsierra/bilo-api/modules/item/fields"
	order "github.com/ricardorsierra/bilo-api/modules/order/fields"
	profile "github.com/ricardorsierra/bilo-api/modules/profile/fields"
	//bank "github.com/ricardorsierra/bilo-api/modules/bank/fields"
	//itemOrder "github.com/ricardorsierra/bilo-api/modules/item_order/fields"
	//itemType "github.com/ricardorsierra/bilo-api/modules/item_type/fields"
	//notification "github.com/ricardorsierra/bilo-api/modules/notification/fields"
	//schedule "github.com/ricardorsierra/bilo-api/modules/schedule/fields"
	//table "github.com/ricardorsierra/bilo-api/modules/table/fields"

	"github.com/graphql-go/graphql"
)

// MutationType is the root mutation
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": post.CreatePostType,
		"createUser": user.CreateUserType,
		"createPlace": place.CreatePlaceType,
		"createBiloCode": biloCode.CreateBiloCodeType,
		"createPrinter": printer.CreatePrinterType,
		"createCreditCard": creditCard.CreateCreditCardType,
		"createItem": item.CreateItemType,
		"createBankAccount": bankAccount.CreateBankAccountType,
		"createProfile": profile.CreateProfileType,
		"createOrder": order.CreateOrderType,
	},
})
