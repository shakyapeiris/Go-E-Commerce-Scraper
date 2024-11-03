package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/shakyapeiris/e-commerce-scraper/config"
	"github.com/shakyapeiris/e-commerce-scraper/logging"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("postgres", config.DB.ConnectionString())
	if err != nil {
		logging.Error("Error connecting to database: %v", err)
		panic(err)
	}
}

func MustBegin() *sqlx.Tx {
	return db.MustBegin()
}

func MustExec(query string, args ...interface{}) sql.Result {
	return db.MustExec(query, args...)
}

func Select(dest interface{}, query string, args ...interface{}) error {
	return db.Select(dest, query, args...)
}

func Get(dest interface{}, query string, args ...interface{}) error {
	return db.Get(dest, query, args...)
}

func NamedExec(query string, arg interface{}) (sql.Result, error) {
	return db.NamedExec(query, arg)
}
