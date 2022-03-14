package project_52weeks

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cheolgyu/pubapi/src/domain"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	usecase domain.ProjectUsecase
}

func NewHandler(r *httprouter.Router, obj_usecase domain.ProjectUsecase) {

	h := Handler{usecase: obj_usecase}
	//chk := CheckHandler{}

	r.GET("/project/day_trading", h.GetDayTradingByPaging)
	r.GET("/project/monthly_peek", h.GetMonthlyPeekByPaging)

}

func (obj *Handler) GetDayTradingByPaging(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")

	////////////////////////////////////////////////////////
	q := r.URL.Query()
	log.Printf("<%s>  params=%s \n", req_id, q)
	paging := domain.ProjectParamsString{}
	paging.Set(q)
	////////////////////////////////////////////////////////

	//ctx := r.Context()
	cmp, err := obj.usecase.GetDayTradingByPaging(context.TODO(), paging)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(cmp)
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)

}

func (obj *Handler) GetMonthlyPeekByPaging(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")

	///////////////////////////////////////////////////
	q := r.URL.Query()
	log.Printf("<%s>  params=%s \n", req_id, q)
	paging := domain.ProjectParamsString{}
	paging.Set(q)
	///////////////////////////////////////////////////

	cmp, err := obj.usecase.GetMonthlyPeekByPaging(context.TODO(), paging)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(cmp)
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)

}
