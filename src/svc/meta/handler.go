package meta

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	usecase domain.MetaUsecase
}

func NewHandler(r *httprouter.Router, cmp_usecase domain.MetaUsecase) {

	h := Handler{usecase: cmp_usecase}
	//chk := CheckHandler{}

	r.GET("/config", h.GetConfig)

}

func (obj *Handler) GetConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")

	//ctx := r.Context()
	cmp, err := obj.usecase.GetConfig(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(w).Encode(cmp)
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~", req_id)

}
