package dao

import (
	"log"
	"strconv"

	"github.com/cheolgyu/stock-read-pub-api/src/model"
	"github.com/jmoiron/sqlx"
)

var SqlDayTrading DayTradingDao

type DayTradingDao struct {
}

func (obj DayTradingDao) Get(req_id string, params model.DatTradingParams) []map[string]interface{} {

	q := `
SELECT *
from public.tb_daily_day_trading
WHERE 1 = 1
	`
	if len(params.Market) > 0 {
		q += `and market in ( `
		for i, v := range params.Market {
			if i > 0 {
				q += ` ,`
			}
			q += ` '` + v + `' `
		}

		q += ` ) `
	}
	if params.Sort != "" {
		q += ` order by  ` + params.Sort + `  ` + params.GetDesc() + ` `
	}
	q += `limit ` + strconv.Itoa(params.Limit) + ` OFFSET ` + strconv.Itoa(params.Offset)

	log.Printf("<%s> query=%s \n", req_id, q)

	var rows *sqlx.Rows
	var err error
	rows, err = DB.Queryx(q)

	if err != nil {
		log.Printf("DayTradingDao:Queryx::error::::<%s>  \n", req_id)
		log.Printf("DayTradingDao:Queryx::error::::<%s> query= \n", q)
		panic(err)
	}

	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		item := make(map[string]interface{})

		err = rows.MapScan(item)

		if err != nil {
			log.Printf("DayTradingDao:MapScan::error::::<%s>  \n", req_id)
			panic(err)
		}
		list = append(list, Decode(item))

	}
	return list
}
