package dao

import (
	"fmt"
	"log"

	"github.com/cheolgyu/stock/backend/api/src/db"
)

var SqlDetail DetailDao

type DetailDao struct {
	db.DB
}

func init() {
	SqlDetail = DetailDao{
		db.DB{},
	}
}

func (obj DetailDao) SelectChart(req_id string, code string, page int) string {

	var db = obj.DB.Conn()
	defer db.Close()

	q := `
SELECT
	json_build_object(
		'date',
		array_agg(tb.d),
		'op',
		array_agg(tb.op),
		'hp',
		array_agg(tb.hp),
		'lp',
		array_agg(tb.lp),
		'cp',
		array_agg(tb.cp),
		'v',
		array_agg(tb.v),
		'fr',
		array_agg(tb.fr)
	) AS item
FROM
	(
		SELECT
			t. *
		FROM
(
				SELECT
					"Date" AS d,
					"OpenPrice" op,
					"HighPrice" hp,
					"LowPrice" lp,
					"ClosePrice" cp,
					"Volume" v,
					"ForeignerBurnoutRate" fr
				FROM
					"price_day"."tb_%s"
				ORDER BY
					"Date" DESC
				LIMIT
					30 offset (% v -1) * 30
			) t
		ORDER BY
			t.d ASC
	) tb
	`
	pq := fmt.Sprintf(q, code, page)
	log.Println(pq)
	var item string
	err := db.QueryRow(pq).Scan(&item)

	if err != nil {
		log.Printf("<%s> error \n", req_id)
		panic(err)
	}

	return item
}

func (obj DetailDao) SelectCompany(req_id string, code string) string {

	var db = obj.DB.Conn()
	defer db.Close()

	q := `
	select
	json_build_object(
		'c',
		row_to_json(c.*),
		'cs',
		row_to_json(cs.*)
		) 
	from listed_company c left join listed_company_state cs on c.short_code = cs.code 
	where c.short_code = '%s'
	limit 1
	`
	pq := fmt.Sprintf(q, code)
	log.Println(pq)
	var item string
	err := db.QueryRow(pq).Scan(&item)

	if err != nil {
		log.Printf("<%s> error \n", req_id)
		panic(err)
	}

	return item
}
