package model

import (
	"strconv"
)

var MPP_AllowSort []string

func init() {
	MPP_AllowSort = append(MPP_AllowSort, "code")
	MPP_AllowSort = append(MPP_AllowSort, "name")
	MPP_AllowSort = append(MPP_AllowSort, "peek")
	MPP_AllowSort = append(MPP_AllowSort, "peek_percent")
}

type MonthlyPeekParms struct {
	Limit  int
	Offset int
	Sort   string
	Desc   bool
	Market []string
}

func (obj *MonthlyPeekParms) GetDesc() string {
	if obj.Desc {
		return "desc"
	} else {
		return "asc"
	}
}

func (obj *MonthlyPeekParms) SetPageRows(page string, rows string) {

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

func (obj *MonthlyPeekParms) SetSortDesc(sort string, in_desc string) {

	desc, err := strconv.ParseBool(in_desc)
	if err != nil {
		desc = true
	}
	obj.Desc = desc
	obj.Sort = MPP_AllowSort[3]

	for _, s := range MPP_AllowSort {
		if s == sort {
			obj.Sort = sort
			break
		}
	}

}

func (obj *MonthlyPeekParms) SetMarket(market string) {

	obj.Market = ParseMarketIn(market)
}
