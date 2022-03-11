package info

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cheolgyu/sf-pub-api/src/domain"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	usecase domain.InfoUsecase
}

func NewHandler(r *httprouter.Router, usecase domain.InfoUsecase) {

	h := Handler{usecase: usecase}
	//chk := CheckHandler{}

	r.GET("/update_time", h.GetUpdateTime)
}

func (obj *Handler) GetUpdateTime(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cmp, err := obj.usecase.GetUpdateTime(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(w).Encode(cmp)
}
