package service

import (
	"github.com/cheolgyu/stock/backend/api/src/dao"
)

func GetDetailChart(req_id string, code string, page int) string {

	return dao.SqlDetail.SelectChart(req_id, code, page)

}

func GetDetailCompany(req_id string, code string) string {

	return dao.SqlDetail.SelectCompany(req_id, code)

}
