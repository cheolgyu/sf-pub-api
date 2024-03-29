package utils

import (
	"log"
	"net/url"
	"strconv"
)

const DefaultRows = 30

type PagingString struct {
	Rows string
	Page string
	Sort string
	Desc string
}

type Paging struct {
	Limit  int
	Offset int
	Sort   string
	Desc   string
}

func (obj *PagingString) Set(query url.Values) {
	obj.Rows = query.Get("rows")
	obj.Page = query.Get("page")
	obj.Sort = query.Get("sort")
	obj.Desc = query.Get("desc")
}

func (obj *PagingString) Valid(sort_column_name map[string][]string, tb_name string) (res Paging, err error) {
	var limit, offset int
	var sort, desc string

	if limit, offset, err = obj.valid_rows_page(); err != nil {
		log.Fatalln(err)
	} else {
		res.Limit = limit
		res.Offset = offset
	}
	if sort_column_name != nil {
		if sort, desc, err = obj.valid_sort_desc(sort_column_name[tb_name]); err != nil {

			res.Sort = sort_column_name[tb_name][0]
			res.Desc = "desc"

			log.Fatalln(err)
		} else {
			res.Sort = sort
			res.Desc = desc
		}
	} else {
		res.Sort = obj.Sort
		if obj.Desc == "true" {
			res.Desc = "desc"
		} else {
			res.Desc = "asc"
		}
	}

	return res, err
}

func (obj *PagingString) valid_rows_page() (limit int, offsest int, err error) {
	p, limit := 1, DefaultRows
	if obj.Page != "" {
		if p, err = strconv.Atoi(obj.Page); err != nil {
			log.Println(err)
		}
	}
	if obj.Rows != "" {
		if limit, err = strconv.Atoi(obj.Rows); err != nil {
			log.Fatalln(err)
		}
	}

	offset := (p - 1) * limit

	return limit, offset, err
}

func (obj *PagingString) valid_sort_desc(sort_column_name []string) (sort string, desc string, err error) {
	var desc_bool bool

	if desc_bool, err = strconv.ParseBool(obj.Desc); err != nil {
		desc = "desc"
		log.Fatalln(err)
	} else {
		if desc_bool {
			desc = "desc"
		} else {
			desc = "asc"
		}
	}

	for i := range sort_column_name {
		if sort_column_name[i] == obj.Sort {
			sort = sort_column_name[i]
			break
		}
	}

	return sort, desc, err
}
