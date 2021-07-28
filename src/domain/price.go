package domain

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Price struct {
}

type PriceUsecase interface {
	GetStockByPaging(ctx context.Context, params PricePagingString) ([]map[string]interface{}, error)
	GetMarketByPaging(ctx context.Context, params PricePagingString) ([]map[string]interface{}, error)
}
type PriceRepository interface {
	GetStockByPaging(ctx context.Context, params PricePaging) ([]map[string]interface{}, error)
	GetMarketByPaging(ctx context.Context, params PricePaging) ([]map[string]interface{}, error)
}

const DefaultRows = 30

type PricePagingString struct {
	Rows   string
	Page   string
	Sort   string
	Desc   string
	Market string
	State  string
	Search string
}

func (obj *PricePagingString) Set(query url.Values) {
	obj.Rows = query.Get("rows")
	obj.Page = query.Get("page")
	obj.Sort = query.Get("sort")
	obj.Desc = query.Get("desc")
	obj.Market = query.Get("market")
	obj.State = query.Get("state")
	obj.Search = query.Get("search")
}

func (obj *PricePagingString) Valid(market_list []Config, column_name map[string][]string, tb_name string) (res PricePaging, err error) {
	var limit, offset int
	var sort, desc, search string
	var market_type []int
	var state map[string]bool

	if limit, offset, err = obj.valid_rows_page(); err != nil {
		log.Fatalln(err)
	} else {
		res.Limit = limit
		res.Offset = offset
	}

	if sort, desc, err = obj.valid_sort_desc(column_name[tb_name]); err != nil {

		res.Sort = column_name[tb_name][0]
		res.Desc = "desc"

		log.Fatalln(err)
	} else {
		res.Sort = sort
		res.Desc = desc
	}

	if search, err = obj.valid_search(); err != nil {
		res.Search = search
		log.Fatalln(err)
	} else {
		res.Search = search
	}

	if market_type, err = obj.valid_marekt_type(market_list); err != nil {
		res.Market = market_type
		log.Println(err)
	} else {
		res.Market = market_type
	}

	if state, err = obj.valid_stock_state(column_name["company_state"]); err != nil {
		res.State = state
		log.Fatalln(err)
	} else {
		res.State = state
	}

	return res, err
}

func (obj *PricePagingString) valid_rows_page() (limit int, offsest int, err error) {
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

func (obj *PricePagingString) valid_sort_desc(column_name []string) (sort string, desc string, err error) {
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

	for i := range column_name {
		if column_name[i] == obj.Sort {
			sort = column_name[i]
			break
		}
	}

	return sort, desc, err
}

func (obj *PricePagingString) valid_search() (search string, err error) {
	search = strings.ReplaceAll(obj.Search, "-", "")
	search = strings.ReplaceAll(search, "'", "")
	search = strings.ReplaceAll(search, ";", "")

	return search, err
}

func (obj *PricePagingString) valid_marekt_type(market_list []Config) (marekt_type []int, err error) {
	if obj.Market != "" {
		inp_marekt_str := obj.Market

		market_str := strings.TrimSpace(inp_marekt_str)
		inp_arr := strings.Split(market_str, ",")

		for i := range market_list {
			for j := range inp_arr {
				var num1 int
				num1, err = strconv.Atoi(inp_arr[j])

				if market_list[i].Id == num1 {
					marekt_type = append(marekt_type, num1)
				}
			}

		}
	} else {
		err = errors.New("market을 확일할 수 없습니다.\n")
	}

	return marekt_type, err
}

func (obj *PricePagingString) valid_stock_state(column_name []string) (res map[string]bool, err error) {
	res = make(map[string]bool)
	inp_str := obj.State
	inp_str = strings.TrimSpace(inp_str)

	inp_arr := strings.Split(inp_str, ",")

	for i := range column_name {
		for j := range inp_arr {
			key_val := strings.Split(inp_arr[j], "::")

			if len(key_val) == 2 {
				if column_name[i] == key_val[0] {
					if v, err := strconv.ParseBool(key_val[1]); err != nil {
						log.Fatalln(err)
					} else {
						res[key_val[0]] = v
					}

				}
			}
		}

	}

	return res, err
}

type PricePaging struct {
	Limit  int
	Offset int
	Sort   string
	Desc   string
	Market []int
	State  map[string]bool
	Search string
}
