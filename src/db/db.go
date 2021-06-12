package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	Info string
}

func (db DB) Conn() *sql.DB {

	conn_db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	return conn_db
}
