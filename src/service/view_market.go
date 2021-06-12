package service

import (
	"github.com/cheolgyu/stock/backend/api/src/dao"
	"github.com/cheolgyu/stock/backend/api/src/model"
)

func GetMarket(req_id string) []model.ViewMarket {
	return dao.SqlMarketDao.Select(req_id)

}
