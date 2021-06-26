package service

import (
	"github.com/cheolgyu/stock-read-pub-api/src/dao"
)

func GetInfo(req_id string) []map[string]interface{} {

	return dao.SqlInfo.Select(req_id)

}
