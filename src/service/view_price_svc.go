package service

import (
	"log"
	"net/http"

	"github.com/cheolgyu/stock/backend/api/src/dao"
	"github.com/cheolgyu/stock/backend/api/src/model"
)

type ViewPrice struct {
}

func GetViewPrice(req_id string, r *http.Request) []model.ViewPrice {

	q := r.URL.Query()
	log.Printf("<%s>  params=%s \n", req_id, q)
	vpp := model.ViewPriceParms{}
	vpp.SetPageRows(q.Get("page"), q.Get("rows"))
	vpp.SetSortDesc(q.Get("sort"), q.Get("desc"))
	vpp.SetEtc(q.Get("market"), q.Get("state"), q.Get("search"))

	return dao.SqlViewPrice.Select(req_id, vpp)

}
