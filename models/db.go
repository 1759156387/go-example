package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DBSTR string

func opendb() (*sql.DB, error) {
	return sql.Open("mysql", DBSTR)
}
