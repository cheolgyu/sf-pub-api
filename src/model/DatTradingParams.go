package model

import (
	"strconv"
)

var DTP_AllowSort []string

func init() {
	DTP_AllowSort = append(DTP_AllowSort, "code")
	DTP_AllowSort = append(DTP_AllowSort, "name")
	DTP_AllowSort = append(DTP_AllowSort, "avg_l2h")
	DTP_AllowSort = append(DTP_AllowSort, "avg_o2c")
	DTP_AllowSort = append(DTP_AllowSort, "std_l2h")
	DTP_AllowSort = append(DTP_AllowSort, "std_o2c")
}

type DatTradingParams struct {
	Term   int
	Limit  int
	Offset int
	Sort   string
	Desc   bool
	Market []string
}

func (obj *DatTradingParams) SetTerm(inp string) {
	p, err := strconv.Atoi(inp)
	if err != nil || p == 0 {
		p = 1
	}
	obj.Term = p
}

func (obj *DatTradingParams) GetDesc() string {
	if obj.Desc {
		return "desc"
	} else {
		return "asc"
	}
}

func (obj *DatTradingParams) SetPageRows(page string, rows string) {

	p, err := strconv.Atoi(page)
	if err != nil || p == 0 {
		p = 1
	}

	limit, err := strconv.Atoi(rows)
	if err != nil || limit == 0 {
		limit = Rows
	}
	offset := (p - 1) * limit

	obj.Limit = limit
	obj.Offset = offset
}

func (obj *DatTradingParams) SetSortDesc(sort string, in_desc string) {

	desc, err := strconv.ParseBool(in_desc)
	if err != nil {
		desc = true
	}
	obj.Desc = desc
	obj.Sort = DTP_AllowSort[2]

	for _, s := range DTP_AllowSort {
		if s == sort {
			obj.Sort = sort
			break
		}
	}

}

func (obj *DatTradingParams) SetMarket(market string) {

	obj.Market = ParseMarketIn(market)
}
