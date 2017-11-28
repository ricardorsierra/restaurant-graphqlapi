package db

import (
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var DB  *sqlx.DB

func init() {
    connStr := os.Getenv("CONNECTION_STRING")
    if connStr=="" {
        connStr = "bilo:GhyY3jGM33Xg1020@tcp(bilo.ciwggvxnyly2.sa-east-1.rds.amazonaws.com:3306)/bilo-backend_production"
    }
    db, err := sqlx.Open("mysql", connStr)
    if err != nil {
        panic(err)
    }
    DB = db
}
