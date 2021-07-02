package dao

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var SqlDayTrading DayTradingDao

type DayTradingDao struct {
}

func (obj DayTradingDao) Get(req_id string, market string, start string, end string) []map[string]interface{} {

	q := `
SELECT C.CODE,
	C.NAME,
	ROUND(AVG(PS.L2H),
		2) AS AVG_L2H,
	ROUND(AVG(PS.O2C),
		2) AS AVG_O2C
FROM COMPANY.STATE S
LEFT JOIN COMPANY.CODE C ON S.CODE = C.CODE
LEFT JOIN COMPANY.DETAIL D ON C.CODE = D.CODE
LEFT JOIN
	(SELECT *
		FROM HIST.PRICE_STOCK
		WHERE 1 = 1
			AND P_DATE BETWEEN %v AND %v
		GROUP BY CODE,
			P_DATE) PS ON C.CODE = PS.CODE
WHERE 1 = 1
	AND S.STOP IS FALSE
	AND D.MARKET = '%v'
GROUP BY C.CODE
ORDER BY AVG_L2H DESC
LIMIT 10
`
	q = fmt.Sprintf(q, start, end, market)
	var rows *sqlx.Rows
	var err error
	rows, err = DB.Queryx(q)

	if err != nil {
		log.Printf("DayTradingDao:Queryx::error::::<%s>  \n", req_id)
		log.Printf("DayTradingDao:Queryx::error::::<%s> query= \n", q)
		panic(err)
	}

	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		item := make(map[string]interface{})

		err = rows.MapScan(item)

		if err != nil {
			log.Printf("DayTradingDao:MapScan::error::::<%s>  \n", req_id)
			panic(err)
		}
		list = append(list, Decode(item))

	}
	return list
}
