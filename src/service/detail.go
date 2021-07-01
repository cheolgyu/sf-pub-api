package service

import (
	"github.com/cheolgyu/stock-read-pub-api/src/dao"
)

func GetDetailChart(req_id string, table_nm string, code string, page int) string {

	return dao.SqlDetail.SelectChart(req_id, table_nm, code, page)

}

func GetDetailCompany(req_id string, code string) string {

	return dao.SqlDetail.SelectCompany(req_id, code)

}
