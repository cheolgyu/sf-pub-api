package domain

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/cheolgyu/sf-pub-api/src/domain/utils"
)

type Project struct {
}

type ProjectUsecase interface {
	GetDayTradingByPaging(ctx context.Context, params ProjectParamsString) ([]map[string]interface{}, error)
	GetMonthlyPeekByPaging(ctx context.Context, params ProjectParamsString) ([]map[string]interface{}, error)
}
type ProjectRepository interface {
	GetDayTradingByPaging(ctx context.Context, params ProjectParams) ([]map[string]interface{}, error)
	GetMonthlyPeekByPaging(ctx context.Context, params ProjectParams) ([]map[string]interface{}, error)
}

type ProjectParamsString struct {
	PagingString utils.PagingString
	Market       string
	Term         string
	UnitType     string
	UnitVal      string
}

type ProjectParams struct {
	Paging   utils.Paging
	Market   []int
	Term     int
	UnitType int
	UnitVal  int
}

func (obj *ProjectParamsString) Set(query url.Values) {
	obj.PagingString.Set(query)
	obj.Market = query.Get("market")
	obj.Term = query.Get("term")
	obj.UnitType = query.Get("unit_type")
	obj.UnitVal = query.Get("unit_val")

}

func (obj *ProjectParamsString) Valid(market_list []Config, sort_column_name map[string][]string, tb_name string) (res ProjectParams, err error) {
	var paging utils.Paging
	var market_type []int
	var term int
	var unit_type int
	var unit_val int

	if paging, err = obj.PagingString.Valid(sort_column_name, tb_name); err != nil {
		log.Fatalln(err)
	}
	res.Paging = paging

	if market_type, err = obj.valid_marekt_type(market_list); err != nil {
		res.Market = market_type
		log.Println(err)
	} else {
		res.Market = market_type
	}

	if term, err = obj.valid_term(); err != nil {
		log.Println(err)
	} else {
		res.Term = term
	}

	if unit_type, unit_val, err = obj.valid_unit_type(); err != nil {
		log.Println(err)
	} else {
		res.UnitType = unit_type
		res.UnitVal = unit_val
	}

	return res, err
}

func (obj *ProjectParamsString) valid_marekt_type(market_list []Config) (marekt_type []int, err error) {
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

func (obj *ProjectParamsString) valid_term() (term int, err error) {
	if obj.Term != "" {
		var num int
		if num, err = strconv.Atoi(obj.Term); err != nil {
			log.Println(err)
		}
		term = num
	}

	return term, err
}

func (obj *ProjectParamsString) valid_unit_type() (unit_type int, unit_val int, err error) {
	if obj.UnitType != "" {
		var num int
		if num, err = strconv.Atoi(obj.UnitType); err != nil {
			log.Println(err)
		}
		unit_type = num

		if num, err = strconv.Atoi(obj.UnitVal); err != nil {
			log.Println(err)
		}
		unit_val = num
	}

	return unit_type, unit_val, err
}
