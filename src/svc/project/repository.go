package project

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/cheolgyu/sf-pub-api/src/domain"
	"github.com/cheolgyu/sf-pub-api/src/domain/utils"
	"github.com/jmoiron/sqlx"
)

type ProjectRepository struct {
	conn *sqlx.DB
}

func NewRepository(Conn *sqlx.DB) domain.ProjectRepository {

	return &ProjectRepository{conn: Conn}
}

func (obj *ProjectRepository) GetDayTradingByPaging(ctx context.Context, params domain.ProjectParams) ([]map[string]interface{}, error) {

	q := "select * from project.func_day_trading(%v, %v, %v, '%s', '%s', ARRAY[ %s])"

	m_str := ""
	for i, v := range params.Market {
		m_str += fmt.Sprintf("%v", v)
		if i+1 < len(params.Market) {
			m_str += ","
		}
	}
	query := fmt.Sprintf(q, params.Term, params.Paging.Limit, params.Paging.Offset, params.Paging.Sort, params.Paging.Desc, m_str)

	log.Printf("query=%s \n", query)

	var rows *sqlx.Rows
	var err error
	rows, err = obj.conn.Queryx(query)

	if err != nil {
		log.Printf("GetDayTradingByPaging:Queryx::error::::<%s> query= \n", q)
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

func (obj *ProjectRepository) GetMonthlyPeekByPaging(ctx context.Context, params domain.ProjectParams) ([]map[string]interface{}, error) {
	log.Println("~~~~~~~ProjectRepository~~~~~~GetMonthlyPeekByPaging~~~~~~~~~~~~~~~~~~~~~~~")
	q := `
SELECT *
from PUBLIC.view_project_trading_volume
WHERE 1 = 1 
	and unit_type=  ` + fmt.Sprintf("%v", params.UnitType) + `
	`
	if params.UnitVal > 0 {
		q += `and  max_unit =   ` + fmt.Sprintf("%v", params.UnitVal)
	}
	if len(params.Market) > 0 {
		q += `and market_type in ( `
		for i, v := range params.Market {
			if i > 0 {
				q += ` ,`
			}
			q += ` ` + fmt.Sprintf("%v", v) + ` `
		}

		q += ` ) `
	}
	if params.Paging.Sort != "" {
		q += ` order by  ` + params.Paging.Sort + `  ` + params.Paging.Desc + ` `
	}
	q += `limit ` + strconv.Itoa(params.Paging.Limit) + ` OFFSET ` + strconv.Itoa(params.Paging.Offset)
	log.Printf("query=%s \n", q)
	rows, err := obj.conn.Queryx(q)
	if err != nil {
		log.Printf("GetMonthlyPeekByPaging:Queryx::error::::<%s> query= \n", q)
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
