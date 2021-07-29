package company

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/cheolgyu/stock-read-pub-api/src/domain/utils"
	"github.com/jmoiron/sqlx"
)

type CompanyRepository struct {
	conn *sqlx.DB
}

func NewRepository(Conn *sqlx.DB) domain.CompanyRepository {

	return &CompanyRepository{conn: Conn}
}

func (obj *CompanyRepository) GetByCode(ctx context.Context, code string) (string, error) {

	q := `
Select JSON_BUILD_OBJECT('c',
	ROW_TO_JSON(C.*),
	'd',
	ROW_TO_JSON(cD.*),
	's',
	ROW_TO_JSON(cS.*),
	 'peek',
	ROW_TO_JSON(mp.*)
					)
from only public.company c 
left join public.company_detail cd on c.code_id = cd.code_id
left join public.company_state cs on c.code_id = cs.code_id
left join public.view_monthly_peek mp on c.code_id = mp.code_id
	WHERE C.CODE = '%s'
	`
	pq := fmt.Sprintf(q, code)
	log.Println(pq)
	var item sql.NullString
	err := obj.conn.QueryRow(pq).Scan(&item)

	if err != nil {
		log.Printf("<%s> error \n", err)
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	if item.Valid {
		return item.String, err
	} else {
		return "", err
	}
}

func (obj *CompanyRepository) GetGraphByCodeID(ctx context.Context, code string, page int) (string, error) {

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
			SELECT dt D,
				concat(SUBSTRING (dt::text,0,5),'-',SUBSTRING (dt::text,5,2),'-',SUBSTRING (dt::text,7,2)) lb,
				CP,
				OP,
				LP,
				HP,
				VOL V,
				FB_RATE FR
			FROM
				hist.price
			WHERE CODE_id =  (select id from meta.code where code ='%s')  
			ORDER BY
				dt DESC
			LIMIT
				30 offset (%v -1) * 30
		) t
		ORDER BY
			t.d ASC
	) tb
	`
	pq := fmt.Sprintf(q, code, page)
	log.Println(pq)
	var item sql.NullString
	err := obj.conn.QueryRow(pq).Scan(&item)

	if err != nil {
		log.Printf("<%s> error \n", err)
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	if item.Valid {
		return item.String, err
	} else {
		return "", err
	}
}

func (obj *CompanyRepository) GetGraphNextLineByCode(ctx context.Context, code string) (string, error) {

	q := `
	select JSON_AGG(json_build_object(price_type ,
		json_build_array(
			json_build_object('x',
			concat(SUBSTRING (x1::text,0,5),'-',SUBSTRING (x1::text,5,2),'-',SUBSTRING (x1::text,7,2)) 
			,'y',y1)
			,json_build_object('x',
			concat(SUBSTRING (x2::text,0,5),'-',SUBSTRING (x2::text,5,2),'-',SUBSTRING (x2::text,7,2))
			,'y',y2)
			,json_build_object('x',
			concat(SUBSTRING (x3::text,0,5),'-',SUBSTRING (x3::text,5,2),'-',SUBSTRING (x3::text,7,2))
			,'y',y3)
		)
	 ))
	from project.tb_line dl  
	where dl.code =  '%s'
	`
	pq := fmt.Sprintf(q, code)
	log.Println(pq)
	var item sql.NullString
	err := obj.conn.QueryRow(pq).Scan(&item)

	if err != nil {
		log.Printf("<%s> error \n", err)
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	if item.Valid {
		return item.String, err
	} else {
		return "", err
	}
}

func (obj *CompanyRepository) GetReboundByPaging(ctx context.Context, params domain.CompanyHisteParams) ([]map[string]interface{}, error) {

	q := `
	SELECT count(*) OVER() AS full_count,*
	from hist.rebound
	WHERE 1 = 1
	and code_id = (select id from meta.code where code= '` + params.Code + `')
	and price_type ='` + fmt.Sprintf("%v", params.Price_type) + `'
	`

	if params.Paging.Sort != "" {
		q += ` order by  ` + params.Paging.Sort + `  ` + params.Paging.Desc + ` `
	}
	q += `limit ` + strconv.Itoa(params.Paging.Limit) + ` OFFSET ` + strconv.Itoa(params.Paging.Offset)

	log.Printf("query=%s \n", q)

	var rows *sqlx.Rows
	var err error
	rows, err = obj.conn.Queryx(q)

	if err != nil {
		log.Printf("GetReboundByPaging:Queryx::error::::<%s> query= \n", q)
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
