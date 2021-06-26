package dao

import (
	"log"
)

var SqlInfo InfoDao

type InfoDao struct {
}

func init() {
	SqlInfo = InfoDao{}
}

func (obj InfoDao) Select(req_id string) []map[string]interface{} {

	q := `

    SELECT
        NAME,
        updated
    FROM
        "info"
	`

	rows, err := DB.Queryx(q)

	if err != nil {
		log.Printf("<%s> error \n", req_id)
		panic(err)
	}

	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {

		item := make(map[string]interface{})
		err = rows.MapScan(item)
		list = append(list, Decode(item))
		if err != nil {
			log.Printf("<%s> error \n", req_id)
			panic(err)
		}

	}

	return list
}
