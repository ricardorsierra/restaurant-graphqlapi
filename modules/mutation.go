package modules

import (
	post "github.com/ricardorsierra/bilo-api/modules/post/fields"
	user "github.com/ricardorsierra/bilo-api/modules/user/fields"


	user "github.com/ricardorsierra/bilo-api/modules/bank/fields"
	user "github.com/ricardorsierra/bilo-api/modules/bank_account/fields"
	user "github.com/ricardorsierra/bilo-api/modules/bilo_code/fields"
	user "github.com/ricardorsierra/bilo-api/modules/credit_card/fields"
	user "github.com/ricardorsierra/bilo-api/modules/item/fields"
	user "github.com/ricardorsierra/bilo-api/modules/item_order/fields"
	user "github.com/ricardorsierra/bilo-api/modules/item_type/fields"
	user "github.com/ricardorsierra/bilo-api/modules/notification/fields"
	user "github.com/ricardorsierra/bilo-api/modules/order/fields"
	user "github.com/ricardorsierra/bilo-api/modules/printer/fields"
	user "github.com/ricardorsierra/bilo-api/modules/profile/fields"
	user "github.com/ricardorsierra/bilo-api/modules/schedule/fields"
	user "github.com/ricardorsierra/bilo-api/modules/table/fields"

	"github.com/graphql-go/graphql"
)

// MutationType is the root mutation
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": post.CreatePostType,
		"createUser": user.CreateUserType,
		"createPlace": place.CreatePlaceType,
	},
})
