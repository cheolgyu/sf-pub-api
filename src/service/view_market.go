package service

import (
	"github.com/cheolgyu/stock-read-http-api/src/dao"
	"github.com/cheolgyu/stock-read-http-api/src/model"
)

func GetMarket(req_id string) []model.ViewMarket {
	return dao.SqlMarketDao.Select(req_id)

}
