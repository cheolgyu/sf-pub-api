package info

import (
	"context"
	"database/sql"
	"log"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/jmoiron/sqlx"
)

type InfoRepository struct {
	conn *sqlx.DB
}

func NewRepository(Conn *sqlx.DB) domain.InfoRepository {

	return &InfoRepository{conn: Conn}
}

func (obj *InfoRepository) GetUpdateTime(ctx context.Context) (string, error) {

	q := `SELECT updated  FROM info where name = 'updated'`
	log.Println(q)
	var item string
	err := obj.conn.QueryRow(q).Scan(&item)

	if err != nil {
		log.Printf("<%s> error \n", err)
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	return item, err
}
