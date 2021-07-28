package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cheolgyu/stock-read-pub-api/src/config"
	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/cheolgyu/stock-read-pub-api/src/service"
	"github.com/julienschmidt/httprouter"
)

type OneStockHandler struct {
	cmp_uc domain.CompanyUsecase
}

func NewOneStockHandler(r *httprouter.Router, _cmp_uc domain.CompanyUsecase) {

	h := OneStockHandler{cmp_uc: _cmp_uc}
	chk := CheckHandler{}

	r.GET("/detail/company/:code", chk.is_market_code(h.GetCompany))
	r.GET("/detail/chart/:code", h.GetChart)
	r.GET("/detail/chartline/:code", h.GetChartLine)

}

func (obj *OneStockHandler) GetCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	// 	setCors(&w)
	//setCors(&w)

	req_code := ps.ByName("code")
	var list string
	market, i := ChkMarketCode(req_code)

	if market {
		e, err := json.Marshal(config.MarketList[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		list = `{ "c":` + string(e) + `}`
	} else {
		ctx := r.Context()
		cmp, err := obj.cmp_uc.GetByCode(ctx, req_code)
		if err != nil {
			log.Fatalln(err)
		}
		json.NewEncoder(w).Encode(cmp)
		log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)
		//list = service.GetDetailCompany(req_id, req_code)
	}

	json.NewEncoder(w).Encode(list)

}

func (obj *OneStockHandler) GetChart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	//setCors(&w)

	req_code := ps.ByName("code")
	q := r.URL.Query()
	page := q.Get("page")

	tbnm := "hist.price_stock"
	market, _ := ChkMarketCode(req_code)
	if market {
		tbnm = "hist.price_market"
	}

	pnum, err := strconv.Atoi(page)
	if err != nil {
		pnum = 1
	}

	list := service.GetDetailChart(req_id, tbnm, req_code, pnum)

	json.NewEncoder(w).Encode(list)

}

func (obj *OneStockHandler) GetChartLine(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	//setCors(&w)

	req_code := ps.ByName("code")

	list := service.GetDetailChartLine(req_id, req_code)

	json.NewEncoder(w).Encode(list)

}
