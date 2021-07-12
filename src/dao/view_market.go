package dao

import (
	"log"

	"github.com/cheolgyu/stock-read-pub-api/src/model"
)

var SqlMarketDao MarketDao

type MarketDao struct {
}

func (obj MarketDao) Select(req_id string, parms model.ViewPriceParams) []map[string]interface{} {

	q := `

    SELECT
        *
    FROM
        "daily_market"
	
	`
	if parms.Sort != "" {
		q += ` order by  ` + parms.Sort + `  ` + parms.GetDesc() + ` `
	}

	rows, err := DB.Queryx(q)

	if err != nil {
		log.Printf("view_market:Queryx::error::::<%s>  \n", req_id)
		log.Printf("view_market:Queryx::error::::<%s> query= \n", q)
		panic(err)
	}

	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		item := make(map[string]interface{})

		err = rows.MapScan(item)
		if err != nil {
			log.Printf("view_market:MapScan::error::::<%s>  \n", req_id)
			panic(err)
		}
		list = append(list, Decode(item))

	}
	return list
}
