package dao

import (
	"log"
)

var SqlMarketDao MarketDao

type MarketDao struct {
}

func (obj MarketDao) Select(req_id string) []map[string]interface{} {

	q := `

    SELECT
        *
    FROM
        "daily_market"
  
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
		if err != nil {
			log.Printf("<%s> error \n", req_id)
			panic(err)
		}
		list = append(list, Decode(item))

	}
	return list
}
