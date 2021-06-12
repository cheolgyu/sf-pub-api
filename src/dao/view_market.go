package dao

import (
	"log"

	"github.com/cheolgyu/stock/backend/api/src/db"
	"github.com/cheolgyu/stock/backend/api/src/model"
)

var SqlMarketDao MarketDao

type MarketDao struct {
	db.DB
}

func init() {
	SqlMarketDao = MarketDao{
		db.DB{},
	}
}

func (obj MarketDao) Select(req_id string) []model.ViewMarket {

	var db = obj.DB.Conn()
	defer db.Close()

	q := `

    SELECT
        *
    FROM
        "view_market_day"
  
	`

	rows, err := db.Query(q)

	if err != nil {
		log.Printf("<%s> error \n", req_id)
		panic(err)
	}

	defer rows.Close()

	var list []model.ViewMarket

	for rows.Next() {
		var item model.ViewMarket

		err = rows.Scan(
			&item.ShortCode, &item.High_date, &item.High_price, &item.Last_date, &item.Last_close_price,
			&item.Contrast_price, &item.Fluctuation_rate, &item.Day_count, &item.High_point_updated_date, &item.Naver_link,
		)

		if err != nil {
			log.Printf("<%s> error \n", req_id)
			panic(err)
		}
		list = append(list, item)

	}
	return list
}
