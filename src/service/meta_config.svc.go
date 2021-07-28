package service

import (
	"log"
	"net/http"

	"github.com/cheolgyu/stock-read-pub-api/src/dao"
	"github.com/cheolgyu/stock-read-pub-api/src/model"
)

func SelectMarketList(req_id string, r *http.Request) []model.Config {

	q := r.URL.Query()
	log.Printf("<%s> SelectMarketList  params=%s \n", req_id, q)

	list := dao.SqlConfigDao.GetMarketList(req_id)

	return list

}
