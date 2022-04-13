package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func New() API {
	var api API
	db, err := sql.Open("mysql", "root:m4m4n5s1@tcp(127.0.0.1:3306)/ass")
	if err != nil {
		panic(err)
	}

	api.Db = db
	return api
}

