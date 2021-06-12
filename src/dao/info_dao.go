package dao

import (
	"log"

	"github.com/cheolgyu/stock/backend/api/src/db"
)

var SqlInfo InfoDao

type InfoDao struct {
	db.DB
}

func init() {
	SqlInfo = InfoDao{
		db.DB{},
	}
}

func (obj InfoDao) Select(req_id string) map[string]string {

	var db = obj.DB.Conn()
	defer db.Close()

	q := `

    SELECT
        NAME,
        fmt_timestamp(updated_date) :: text AS updated_date
    FROM
        "info"
    WHERE
        NAME LIKE 'daily_company%'
        OR NAME LIKE 'daily_price_day%'
        OR NAME LIKE 'daily_high%'
        OR NAME LIKE 'daily_market%'
	`

	rows, err := db.Query(q)

	if err != nil {
		log.Printf("<%s> error \n", req_id)
		panic(err)
	}

	defer rows.Close()

	//var list []interface{}
	item := make(map[string]string)

	for rows.Next() {

		var nm string
		var val string
		err = rows.Scan(&nm, &val)
		item[nm] = val
		if err != nil {
			log.Printf("<%s> error \n", req_id)
			panic(err)
		}

	}

	return item
}
