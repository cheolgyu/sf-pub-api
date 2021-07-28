package price

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/jmoiron/sqlx"
)

func Decode(item map[string]interface{}) map[string]interface{} {
	for k, v := range item {
		if b, ok := v.([]byte); ok {
			item[k] = string(b)
		}
		//item[k] = string(v) //v.decode("base64")
	}
	return item
}

type PriceRepository struct {
	conn *sqlx.DB
}

func NewRepository(Conn *sqlx.DB) domain.PriceRepository {

	return &PriceRepository{conn: Conn}
}

func (obj *PriceRepository) GetStockByPaging(ctx context.Context, params domain.PricePaging) ([]map[string]interface{}, error) {

	q := `SELECT count(*) OVER() AS full_count,* FROM  view_stock `
	q += ` where  1=1 `
	if params.Search != "" {
		q += ` and  name  like '%` + params.Search + `%' `
	}
	if len(params.State) > 0 {
		q += ` and (  `
		i := 0
		for k, v := range params.State {
			if i > 0 {
				q += ` and `
			}
			q += fmt.Sprintf(" %s is  %t ", k, v)
			i++
		}
		q += ` ) `
	}

	if len(params.Market) > 0 {
		q += `and market_type in ( `
		for i, v := range params.Market {
			if i > 0 {
				q += ` ,`
			}
			q += fmt.Sprintf(`'%v'`, v)
		}

		q += ` ) `
	}

	if params.Sort != "" {
		q += ` order by  ` + params.Sort + `  ` + params.Desc + ` `
	}
	q += `limit ` + strconv.Itoa(params.Limit) + ` OFFSET ` + strconv.Itoa(params.Offset)

	log.Printf("query=%s \n", q)

	var rows *sqlx.Rows
	var err error
	rows, err = obj.conn.Queryx(q)

	if err != nil {
		log.Printf("view_price:Queryx::error::::<%s> query= \n", q)
		panic(err)
	}

	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		item := make(map[string]interface{})

		err = rows.MapScan(item)

		if err != nil {
			panic(err)
		}
		list = append(list, Decode(item))
	}

	return list, err
}

func (obj *PriceRepository) GetMarketByPaging(ctx context.Context, params domain.PricePaging) ([]map[string]interface{}, error) {

	q := `  SELECT  *  FROM  "view_market" 	`
	if params.Sort != "" {
		q += ` order by  ` + params.Sort + `  ` + params.Desc + ` `
	}
	rows, err := obj.conn.Queryx(q)
	if err != nil {
		log.Printf("view_market:Queryx::error::::<%s> query= \n", q)
		panic(err)
	}

	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		item := make(map[string]interface{})

		err = rows.MapScan(item)

		if err != nil {
			panic(err)
		}
		list = append(list, Decode(item))
	}

	return list, err
}
