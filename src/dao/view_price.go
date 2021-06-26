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
		q += ` and  name like $1 `
	}

	if parms.State != "" {
		_q := ` and %s is true `

		q += fmt.Sprintf(_q, parms.State)
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
	if parms.Search != "" {

		//rows, err = DB.Queryx(q, "%"+parms.Search+"%")
		q += "%" + parms.Search + "%"
		rows, err = DB.Queryx(q)
	} else {
		rows, err = DB.Queryx(q)
	}

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
