package service

import (
	"log"
	"net/http"

	"github.com/cheolgyu/stock-read-pub-api/src/dao"
	"github.com/cheolgyu/stock-read-pub-api/src/model"
)

func GetMonthlyPeek(req_id string, r *http.Request) []map[string]interface{} {

	q := r.URL.Query()
	log.Printf("<%s> DayTradingSvc  params=%s \n", req_id, q)

	var list []map[string]interface{}
	params := model.MonthlyPeekParams{}
	params.SetMarket(q.Get("market"))
	params.SetSortDesc(q.Get("sort"), q.Get("desc"))
	params.SetPageRows(q.Get("page"), q.Get("rows"))

	list = dao.SqlMonthlyPeek.Get(req_id, params)

	return list

}
