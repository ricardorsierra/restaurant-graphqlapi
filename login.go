package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	setupServer()
}


func setupMux() *http.ServeMux {
	mux := http.NewServeMux()

	// graphql Handler
	graphqlHandler := http.HandlerFunc(graphqlHandlerFunc)

	// login Handler
	mux.HandleFunc("/login", loginFunc)

	// add in addContext middlware
	mux.Handle("/graphql", requireAuth(graphqlHandler))

	return mux
}

func setupServer() {
	http.ListenAndServe(":8080", setupMux())
}

// graphqlHandlerFunc creates the graphql handler
func graphqlHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// get query
	opts := handler.NewRequestOptions(r)

	// execute graphql query
	params := graphql.Params{
		Schema:         Schema, // defined in another file
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        r.Context(), // pass http.Request.Context() to our graphql object
	}
	result := graphql.Do(params)

	// output JSON
	var buff []byte
	w.WriteHeader(http.StatusOK)
	if prettyPrintGraphQL {
		buff, _ = json.MarshalIndent(result, "", "\t")
	} else {
		buff, _ = json.Marshal(result)
	}
	w.Write(buff)
}

// type definition for our claims
type Claims struct {
	UserID  uint64 `json:"userID"`
	IsAdmin bool   `json:"isAdmin"`
	jwt.StandardClaims
}

// secret string for signing requests
var jwtSecret = []byte("secret") // make sure you change this to something secure

// key type is not exported to prevent collisions with context keys defined in
// other packages.
type key int

// userAuthKey is the context key for our added struct.  Its value of zero is
// arbitrary.  If this package defined other context keys, they would have
// different integer values.
const userAuthKey key = 0

// validate JWT submitted via Authorization Header and set context claims
func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extract jwt
		authorizationHeader := r.Header.Get("Authorization")
		authRegex, _ := regexp.Compile("(?:Bearer *)([^ ]+)(?: *)")
		authRegexMatches := authRegex.FindStringSubmatch(authorizationHeader)
		if len(authRegexMatches) != 2 {
			// didn't match valid Authorization header pattern
			httpError(w, "not authorized", http.StatusUnauthorized)
			return
		}
		jwtToken := authRegexMatches[1]

		// parse tokentoken
		token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return jwtSecret, nil
			})
		if err != nil {
			httpError(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// extract claims
		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			httpError(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// load userID & isAdmin into context
		authContext := struct {
			UserID  uint64 `json:"userId"`
			IsAdmin bool   `json:"isAdmin"`
		}{
			claims.UserID,
			claims.IsAdmin,
		}
		ctx := context.WithValue(r.Context(), userAuthKey, authContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// loginFunc confirms login credentials and creates JWT if valid
func loginFunc(w http.ResponseWriter, req *http.Request) {
	// get username & password
	decoder := json.NewDecoder(req.Body)
	requestBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// confirmLogin is up to you to define
	user, err := confirmLogin(requestBody)
	if err != nil {
		http.Error(w, "invalid login", http.StatusUnauthorized)
		return
	}

	//generate token
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claims := Claims{
		user.ID,
		user.IsAdmin,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:8080",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(jwtSecret)

	//output token
	tokenResponse := struct {
		Token string `json:"token"`
	}{signedToken}
	json.NewEncoder(w).Encode(tokenResponse)
}