package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/graphql-go/graphql"

	// Database and yours Drives
	"database/sql"
	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

var db *sql.DB

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("mysql", "bilo:GhyY3jGM33Xg1020@tcp(bilo.ciwggvxnyly2.sa-east-1.rds.amazonaws.com:3306)/bilo-backend_production")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler(schema))
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
