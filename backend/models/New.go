package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func New() API {
	var api API
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}

	api.Db = db
	return api
}

