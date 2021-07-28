package dao

import (
	"fmt"
	"log"

	"github.com/cheolgyu/stock-read-pub-api/src/model"
	"github.com/jmoiron/sqlx"
)

var SqlDayTrading DayTradingDao

type DayTradingDao struct {
}

func (obj DayTradingDao) Get(req_id string, params model.DatTradingParams) []map[string]interface{} {
	// select * from project.func_day_trading(10,10,1,'avg_l2h','desc',ARRAY[	7,9]);
	q := "select * from project.func_day_trading(%v, %v, %v, %v, %s, ARRAY[ %s])"
	query := fmt.Sprintf(q, params.Term, params.Limit, params.Offset, params.Sort, params.GetDesc(), params.Market)

	log.Printf("<%s> query=%s \n", req_id, query)

	var rows *sqlx.Rows
	var err error
	rows, err = DB.Queryx(query)

	if err != nil {
		log.Printf("DayTradingDao:Queryx::error::::<%s>  \n", req_id)
		log.Printf("DayTradingDao:Queryx::error::::<%s> query= \n", query)
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
