package dao

import (
	"log"
	"strconv"

	"github.com/cheolgyu/stock-read-pub-api/src/model"
	"github.com/jmoiron/sqlx"
)

var SqlHistPriceBoundDao HistPriceBoundDao

type HistPriceBoundDao struct {
}

func (obj HistPriceBoundDao) SelectList(req_id string, params model.HistPriceParams) []map[string]interface{} {

	q := `
	SELECT count(*) OVER() AS full_count,*
	from ` + params.TbName + `
	WHERE 1 = 1
	and code ='` + params.Code + `'
	and g_type ='` + params.G_type + `'
	`

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
