package service

import (
	"log"
	"net/http"

	"github.com/cheolgyu/stock-read-pub-api/src/dao"
	"github.com/cheolgyu/stock-read-pub-api/src/model"
)

func GetMarket(req_id string, r *http.Request) []map[string]interface{} {
	q := r.URL.Query()
	log.Printf("<%s>  market-params=%s \n", req_id, q)
	vpp := model.ViewPriceParams{}
	vpp.SetSortDesc(q.Get("sort"), q.Get("desc"))

	return dao.SqlMarketDao.Select(req_id, vpp)

}
