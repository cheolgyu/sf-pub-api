package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Conn() *sqlx.DB {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Panic("Error loading .env file")
	}
	DB_URL := os.Getenv("DB_URL")
	log.Println("============================")
	log.Println("============================")
	log.Println("============================")
	log.Println("============================")
	log.Println("============================")
	log.Println(DB_URL)

	conn_db, err := sqlx.Open("postgres", DB_URL)
	if err != nil {
		log.Println("커넥션 오류:", err)
		panic(err)
	}

	return conn_db
}
