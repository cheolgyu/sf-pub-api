package service

import (
	"github.com/cheolgyu/stock-read-pub-api/src/dao"
)

func GetMarket(req_id string) []map[string]interface{} {
	return dao.SqlMarketDao.Select(req_id)

}
