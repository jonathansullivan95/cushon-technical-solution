package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Startup() error {
	var err error
	db, err = sql.Open("mysql", "tester:pa55word@tcp(localhost:3306)/test")
	if err != nil {
		return err
	}

	return nil
}
