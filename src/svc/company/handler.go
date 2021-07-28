package company

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	usecase domain.CompanyUsecase
}

func NewCompanyHandler(r *httprouter.Router, cmp_usecase domain.CompanyUsecase) {

	h := Handler{usecase: cmp_usecase}
	//chk := CheckHandler{}

	r.GET("/company/:code", h.GetCompany)
	r.GET("/company_chart/:code", h.GetGraphByCodeID)
	r.GET("/company_chart_next/:code", h.GetGraphNextLineByCode)

}

func (obj *Handler) GetCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")

	req_code := ps.ByName("code")

	//ctx := r.Context()
	cmp, err := obj.usecase.GetByCode(context.TODO(), req_code)
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(w).Encode(cmp)
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)

}

func (obj *Handler) GetGraphByCodeID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")

	q := r.URL.Query()
	page := q.Get("page")

	pnum, err := strconv.Atoi(page)
	if err != nil {
		pnum = 1
	}

	cmp, err := obj.usecase.GetGraphByCodeID(context.TODO(), ps.ByName("code"), pnum)
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(w).Encode(cmp)
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)

}

func (obj *Handler) GetGraphNextLineByCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	//setCors(&w)

	req_code := ps.ByName("code")

	cmp, err := obj.usecase.GetGraphNextLineByCode(context.TODO(), req_code)
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(w).Encode(cmp)
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)

}