package service

import (
	"log"
	"net/http"

	"github.com/cheolgyu/stock-read-pub-api/src/dao"
	"github.com/cheolgyu/stock-read-pub-api/src/model"
)

func GetHistPriceBound(req_id string, r *http.Request, tbnm string, code string) []map[string]interface{} {

	q := r.URL.Query()
	log.Printf("<%s>  params=%s \n", req_id, q)
	params := model.HistPriceParams{}
	params.TbName = tbnm
	params.Code = code

	params.SetG_type(q.Get("g_type"))
	params.SetPageRows(q.Get("page"), q.Get("rows"))
	params.SetSortDesc(q.Get("sort"), q.Get("desc"))

	return dao.SqlHistPriceBoundDao.SelectList(req_id, params)

}
