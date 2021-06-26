package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Conn() *sqlx.DB {

	conn_db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=example dbname=dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	return conn_db
}
