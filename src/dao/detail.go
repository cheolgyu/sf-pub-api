package dao

import (
	"fmt"
	"log"
)

var SqlDetail DetailDao

type DetailDao struct {
}

func (obj DetailDao) SelectChart(req_id string, code string, page int) string {

	q := `
SELECT
	json_build_object(
		'date',
		array_agg(tb.lb),
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
			SELECT P_DATE D,
				concat(SUBSTRING (P_DATE::text,0,5),'-',SUBSTRING (P_DATE::text,5,2),'-',SUBSTRING (P_DATE::text,7,2)) lb,
				CP,
				OP,
				LP,
				HP,
				VOL V,
				FB_RATE FR
			FROM
				hist.price_stock
			WHERE CODE = '%s'
			ORDER BY
				p_date DESC
			LIMIT
				30 offset (%v -1) * 30
		) t
		ORDER BY
			t.d ASC
	) tb
	`
	pq := fmt.Sprintf(q, code, page)
	log.Println(pq)
	var item string
	err := DB.QueryRow(pq).Scan(&item)

	if err != nil {
		log.Printf("<%s> error \n", req_id)
		panic(err)
	}

	return item
}

func (obj DetailDao) SelectCompany(req_id string, code string) string {

	q := `
	SELECT JSON_BUILD_OBJECT('c',
		ROW_TO_JSON(C.*),
		'd',
		ROW_TO_JSON(D.*),
		's',
		ROW_TO_JSON(S.*))
	FROM COMPANY.CODE C
	LEFT JOIN COMPANY.DETAIL D ON C.CODE = D.CODE
	LEFT JOIN COMPANY.STATE S ON C.CODE = S.CODE
	WHERE C.CODE = '%s'
	`
	pq := fmt.Sprintf(q, code)
	log.Println(pq)
	var item string
	err := DB.QueryRow(pq).Scan(&item)

	if err != nil {
		log.Printf("<%s> error \n", req_id)
		panic(err)
	}

	return item
}
