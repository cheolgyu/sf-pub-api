package model

import (
	"strconv"
)

var HPP_AllowSort []string

func init() {
	HPP_AllowSort = append(HPP_AllowSort, "g_type")
	HPP_AllowSort = append(HPP_AllowSort, "x1")
	HPP_AllowSort = append(HPP_AllowSort, "y1")
	HPP_AllowSort = append(HPP_AllowSort, "x2")
	HPP_AllowSort = append(HPP_AllowSort, "y2")
	HPP_AllowSort = append(HPP_AllowSort, "x_tick")
	HPP_AllowSort = append(HPP_AllowSort, "y_minus")
	HPP_AllowSort = append(HPP_AllowSort, "y_percent")
}

type HistPriceParams struct {
	TbName string
	Code   string
	G_type string
	Limit  int
	Offset int
	Sort   string
	Desc   bool
}

func (obj *HistPriceParams) SetG_type(g_type string) {
	res := AllowG_Type[0]
	for _, v := range AllowG_Type {
		if v == g_type {
			res = g_type
			break
		}
	}
	obj.G_type = res
}

func (obj *HistPriceParams) GetDesc() string {
	if obj.Desc {
		return "desc"
	} else {
		return "asc"
	}
}

func (obj *HistPriceParams) SetPageRows(page string, rows string) {

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

func (obj *HistPriceParams) SetSortDesc(sort string, in_desc string) {

	desc, err := strconv.ParseBool(in_desc)
	if err != nil {
		desc = true
	}
	obj.Desc = desc
	obj.Sort = HPP_AllowSort[1]

	for _, s := range HPP_AllowSort {
		if s == sort {
			obj.Sort = sort
			break
		}
	}

}
