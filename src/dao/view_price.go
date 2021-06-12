package dao

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/cheolgyu/stock/backend/api/src/db"
	"github.com/cheolgyu/stock/backend/api/src/model"
)

var SqlViewPrice ViewPrice

type ViewPrice struct {
	db.DB
}

func init() {
	SqlViewPrice = ViewPrice{
		db.DB{},
	}
}

func (obj ViewPrice) Select(req_id string, parms model.ViewPriceParms) []model.ViewPrice {

	var db = obj.DB.Conn()
	defer db.Close()

	q := `SELECT count(*) OVER() AS full_count,* FROM  view_price_day `

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

	var rows *sql.Rows
	var err error
	if parms.Search != "" {

		rows, err = db.Query(q, "%"+parms.Search+"%")
	} else {
		rows, err = db.Query(q)
	}

	if err != nil {
		log.Printf("<%s> error \n", req_id)
		panic(err)
	}

	defer rows.Close()

	var list []model.ViewPrice

	for rows.Next() {
		var item model.ViewPrice

		err = rows.Scan(
			&item.Full_count,
			&item.Code, &item.Name, &item.Market, &item.High_date, &item.High_price,
			&item.Last_close_price, &item.Contrast_price, &item.Fmt_high_date, &item.Fmt_high_price, &item.Fmt_last_date,
			&item.Fmt_last_close_price, &item.Fmt_contrast_price, &item.Fluctuation_rate, &item.Day_count, &item.Updated_date_high_point,
			&item.Naver_link, &item.Stop, &item.Clear, &item.Managed, &item.Ventilation,
			&item.Unfaithful, &item.Low_liquidity, &item.Lack_listed, &item.Overheated, &item.Caution,
			&item.Warning, &item.Risk, &item.Updated_date_company_state,
		)

		if err != nil {
			log.Printf("<%s> error \n", req_id)
			panic(err)
		}
		list = append(list, item)

	}
	return list
}
