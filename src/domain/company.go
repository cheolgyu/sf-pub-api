package domain

import (
	"context"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/cheolgyu/stock-read-pub-api/src/domain/utils"
)

type Company struct {
}

type CompanyUsecase interface {
	GetByCode(ctx context.Context, code string) (string, error)
	GetGraphByCodeID(ctx context.Context, code string, page int) (string, error)
	GetGraphNextLineByCode(ctx context.Context, code string) (string, error)

	GetHistByPaging(ctx context.Context, params CompanyHisteParamsStr) ([]map[string]interface{}, error)
}

type CompanyRepository interface {
	GetByCode(ctx context.Context, code string) (string, error)
	GetGraphByCodeID(ctx context.Context, code string, page int) (string, error)
	GetGraphNextLineByCode(ctx context.Context, code string) (string, error)

	GetHistByPaging(ctx context.Context, params CompanyHisteParams) ([]map[string]interface{}, error)
	//Fetch(ctx context.Context, cursor string, num int64) (res []Company, nextCursor string, err error)
	//GetByID(ctx context.Context, id int64) (Company, error)
	//GetByTitle(ctx context.Context, title string) (Company, error)
	//Update(ctx context.Context, ar *Company) error
	//Store(ctx context.Context, a *Company) error
	//Delete(ctx context.Context, id int64) error
}

type CompanyHisteParamsStr struct {
	PagingStr  utils.PagingStr
	Code       string
	Price_type string
}

type CompanyHisteParams struct {
	Paging     utils.Paging
	Code       string
	Price_type int
}

func (obj *CompanyHisteParamsStr) Set(query url.Values) {
	obj.PagingStr.Set(query)
	obj.Code = query.Get("code")
	obj.Price_type = query.Get("price_type")
}

func (obj *CompanyHisteParamsStr) Valid(price_type_list []Config, sort_column_name map[string][]string, tb_name string) (res CompanyHisteParams, err error) {
	var paging utils.Paging
	var code string
	var price_type int

	if paging, err = obj.PagingStr.Valid(sort_column_name, tb_name); err != nil {
		log.Fatalln(err)
	}
	res.Paging = paging

	if code, err = obj.valid_code(); err != nil {
		res.Code = code
		log.Fatalln(err)
	} else {
		res.Code = code
	}

	if price_type, err = obj.valid_price_type(price_type_list); err != nil {
		res.Price_type = price_type
		log.Fatalln(err)
	} else {
		res.Price_type = price_type
	}

	return res, err
}

func (obj *CompanyHisteParamsStr) valid_price_type(price_type_list []Config) (price_type int, err error) {

	if obj.Price_type != "" {
		var num int
		if num, err = strconv.Atoi(obj.Price_type); err != nil {
			log.Println(err)
		} else {
			for i := range price_type_list {
				if price_type_list[i].Id == num {
					price_type = price_type_list[i].Id
					break
				}
			}
		}
	}

	return price_type, err
}

func (obj *CompanyHisteParamsStr) valid_code() (code string, err error) {

	if obj.Code != "" {
		code = strings.ReplaceAll(obj.Code, "-", "")
		code = strings.ReplaceAll(code, "'", "")
		code = strings.ReplaceAll(code, ";", "")

		return code, err
	}

	return code, err
}
