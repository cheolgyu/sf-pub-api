package price

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	usecase domain.PriceUsecase
}

func NewHandler(r *httprouter.Router, cmp_usecase domain.PriceUsecase) {

	h := Handler{usecase: cmp_usecase}
	//chk := CheckHandler{}

	r.GET("/price/stock", h.GetStockByPaging)
	r.GET("/price/market", h.GetMarketByPaging)

}

func (obj *Handler) GetStockByPaging(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")

	////////////////////////////////////////////////////////
	q := r.URL.Query()
	log.Printf("<%s>  params=%s \n", req_id, q)
	paging := domain.PricePagingString{}
	paging.Set(q)
	////////////////////////////////////////////////////////

	//ctx := r.Context()
	cmp, err := obj.usecase.GetStockByPaging(context.TODO(), paging)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(cmp)
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)

}

func (obj *Handler) GetMarketByPaging(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")

	///////////////////////////////////////////////////
	q := r.URL.Query()
	log.Printf("<%s>  params=%s \n", req_id, q)
	paging := domain.PricePagingString{}
	paging.Set(q)

	//return dao.SqlMarketDao.Select(req_id, vpp)
	///////////////////////////////////////////////////

	cmp, err := obj.usecase.GetMarketByPaging(context.TODO(), paging)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(cmp)
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)

}
