package domain

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/cheolgyu/stock-read-pub-api/src/domain/utils"
)

type Price struct {
}

type PriceUsecase interface {
	GetStockByPaging(ctx context.Context, params PriceParamsString) ([]map[string]interface{}, error)
	GetMarketByPaging(ctx context.Context, params PriceParamsString) ([]map[string]interface{}, error)
}
type PriceRepository interface {
	GetStockByPaging(ctx context.Context, params PriceParams) ([]map[string]interface{}, error)
	GetMarketByPaging(ctx context.Context, params PriceParams) ([]map[string]interface{}, error)
}

type PriceParamsString struct {
	PagingStr utils.PagingStr
	Market    string
	State     string
	Search    string
}

type PriceParams struct {
	Paging utils.Paging
	Market []int
	State  map[string]bool
	Search string
}

func (obj *PriceParamsString) Set(query url.Values) {
	obj.PagingStr.Set(query)
	obj.Market = query.Get("market")
	obj.State = query.Get("state")
	obj.Search = query.Get("search")
}

func (obj *PriceParamsString) Valid(market_list []Config, sort_column_name map[string][]string, tb_name string) (res PriceParams, err error) {
	var paging utils.Paging
	var search string
	var market_type []int
	var state map[string]bool

	if paging, err = obj.PagingStr.Valid(sort_column_name, tb_name); err != nil {
		log.Fatalln(err)
	}
	res.Paging = paging

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

	if state, err = obj.valid_stock_state(sort_column_name["company_state"]); err != nil {
		res.State = state
		log.Fatalln(err)
	} else {
		res.State = state
	}

	return res, err
}

func (obj *PriceParamsString) valid_search() (search string, err error) {
	search = strings.ReplaceAll(obj.Search, "-", "")
	search = strings.ReplaceAll(search, "'", "")
	search = strings.ReplaceAll(search, ";", "")

	return search, err
}

func (obj *PriceParamsString) valid_marekt_type(market_list []Config) (marekt_type []int, err error) {
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

func (obj *PriceParamsString) valid_stock_state(sort_column_name []string) (res map[string]bool, err error) {
	res = make(map[string]bool)
	inp_str := obj.State
	inp_str = strings.TrimSpace(inp_str)

	inp_arr := strings.Split(inp_str, ",")

	for i := range sort_column_name {
		for j := range inp_arr {
			key_val := strings.Split(inp_arr[j], "::")

			if len(key_val) == 2 {
				if sort_column_name[i] == key_val[0] {
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
