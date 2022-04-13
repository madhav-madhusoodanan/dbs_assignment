package models

import (
	"database/sql"
)

type API struct {
	Db *sql.DB
}

