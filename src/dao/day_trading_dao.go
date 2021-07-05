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

func (obj DayTradingDao) Get(req_id string, parms model.DatTradingParms) []map[string]interface{} {

	q := `
SELECT *
from public.tb_daily_day_trading
WHERE 1 = 1
	`
	if len(parms.Market) > 0 {
		q += `and market in ( `
		for i, v := range parms.Market {
			if i > 0 {
				q += ` ,`
			}
			q += ` '` + v + `' `
		}

		q += ` ) `
	}
	if parms.Sort != "" {
		q += ` order by  ` + parms.Sort + `  ` + parms.GetDesc() + ` `
	}
	q += `limit ` + strconv.Itoa(parms.Limit) + ` OFFSET ` + strconv.Itoa(parms.Offset)

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
