package dao

import (
	"log"
	"strconv"

	"github.com/cheolgyu/stock-read-pub-api/src/model"
	"github.com/jmoiron/sqlx"
)

var SqlMonthlyPeek MonthlyPeekDao

type MonthlyPeekDao struct {
}

func (obj MonthlyPeekDao) Get(req_id string, params model.MonthlyPeekParams) []map[string]interface{} {

	q := `
SELECT *
from public.daily_monthly_peek
WHERE 1 = 1 and stop is false 
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
		log.Printf("MonthlyPeekDao:Queryx::error::::<%s>  \n", req_id)
		log.Printf("MonthlyPeekDao:Queryx::error::::<%s> query= \n", q)
		panic(err)
	}

	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		item := make(map[string]interface{})

		err = rows.MapScan(item)

		if err != nil {
			log.Printf("MonthlyPeekDao:MapScan::error::::<%s>  \n", req_id)
			panic(err)
		}
		list = append(list, Decode(item))

	}
	return list
}
