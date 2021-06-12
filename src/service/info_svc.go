package service

import (
	"github.com/cheolgyu/stock-read-http-api/src/dao"
)

func GetInfo(req_id string) map[string]string {

	return dao.SqlInfo.Select(req_id)

}
