package price

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/cheolgyu/sf-pub-api/src/domain"
	"github.com/cheolgyu/sf-pub-api/src/domain/utils"
	"github.com/jmoiron/sqlx"
)

type PriceRepository struct {
	conn *sqlx.DB
}

func NewRepository(Conn *sqlx.DB) domain.PriceRepository {

	return &PriceRepository{conn: Conn}
}

func (obj *PriceRepository) GetStockByPaging(ctx context.Context, params domain.PriceParams) ([]map[string]interface{}, error) {

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

	if params.Paging.Sort != "" {
		q += ` order by  ` + params.Paging.Sort + `  ` + params.Paging.Desc + ` `
	}
	q += `limit ` + strconv.Itoa(params.Paging.Limit) + ` OFFSET ` + strconv.Itoa(params.Paging.Offset)

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
		list = append(list, utils.Decode(item))
	}

	return list, err
}

func (obj *PriceRepository) GetMarketByPaging(ctx context.Context, params domain.PriceParams) ([]map[string]interface{}, error) {

	q := `  SELECT  *  FROM  "view_market" 	`
	if params.Paging.Sort != "" {
		q += ` order by  ` + params.Paging.Sort + `  ` + params.Paging.Desc + ` `
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
		list = append(list, utils.Decode(item))
	}

	return list, err
}
