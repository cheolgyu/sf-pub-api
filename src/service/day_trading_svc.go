package service

import (
	"log"
	"net/http"

	"github.com/cheolgyu/stock-read-pub-api/src/dao"
	"github.com/cheolgyu/stock-read-pub-api/src/model"
)

type DayTradingSvc struct {
}

func GetDayTrading(req_id string, r *http.Request) []map[string]interface{} {

	q := r.URL.Query()
	log.Printf("<%s> DayTradingSvc  params=%s \n", req_id, q)
	qm := q.Get("market")
	qs := q.Get("start")
	qe := q.Get("end")

	var list []map[string]interface{}

	if model.ChkMarket(qm) && model.ChkDate(qs) && model.ChkDate(qe) {
		list = dao.SqlDayTrading.Get(req_id, qm, qs, qe)
	} else {
		log.Printf("<%s> DayTradingSvc:::::valid params 실패 \n", req_id)
	}

	return list

}
