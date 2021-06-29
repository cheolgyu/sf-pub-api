package dao

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cheolgyu/stock-read-pub-api/src/model"
	"github.com/jmoiron/sqlx"
)

var SqlViewPrice ViewPrice

type ViewPrice struct {
}

func (obj ViewPrice) Select(req_id string, parms model.ViewPriceParms) []map[string]interface{} {

	q := `SELECT count(*) OVER() AS full_count,* FROM  daily_stock `

	q += ` where  1=1 `

	if parms.Search != "" {
		q += ` and  name like '%` + parms.Search + `%' `
	}
	if len(parms.State) > 0 {
		q += ` and (  `
		i := 0
		for k, v := range parms.State {
			if i > 0 {
				q += ` and `
			}
			q += fmt.Sprintf(" %s is  %t ", k, v)
			i++
		}
		q += ` ) `
	}

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
		log.Printf("view_price:Queryx::error::::<%s>  \n", req_id)
		log.Printf("view_price:Queryx::error::::<%s> query= \n", q)
		panic(err)
	}

	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		item := make(map[string]interface{})

		err = rows.MapScan(item)

		if err != nil {
			log.Printf("view_price:MapScan::error::::<%s>  \n", req_id)
			panic(err)
		}
		list = append(list, Decode(item))

	}
	return list
}
