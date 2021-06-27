package model

import (
	"database/sql"
	"strconv"
	"strings"
)

const Rows = 10

var AllowSort [26]string
var AllowMarket [3]string
var AllowState [11]string

func init() {
	var allSort [26]string
	allSort[0] = "code"
	allSort[1] = "name"
	allSort[2] = "cp_x1"
	allSort[3] = "cp_y1"
	allSort[4] = "cp_x2"
	allSort[5] = "cp_y2"
	allSort[6] = "cp_y_percent"
	allSort[7] = "cp_x_tick"

	allSort[8] = "op_x1"
	allSort[9] = "op_y1"
	allSort[10] = "op_x2"
	allSort[11] = "op_y2"
	allSort[12] = "op_y_percent"
	allSort[13] = "op_x_tick"

	allSort[14] = "lp_x1"
	allSort[15] = "lp_y1"
	allSort[16] = "lp_x2"
	allSort[17] = "lp_y2"
	allSort[18] = "lp_y_percent"
	allSort[19] = "lp_x_tick"

	allSort[20] = "hp_x1"
	allSort[21] = "hp_y1"
	allSort[22] = "hp_x2"
	allSort[23] = "hp_y2"
	allSort[24] = "hp_y_percent"
	allSort[25] = "hp_x_tick"
	AllowSort = allSort

	var allMarket [3]string
	allMarket[0] = "KOSPI"
	allMarket[1] = "KOSDAQ"
	allMarket[2] = "KONEX"

	AllowMarket = allMarket

	var allowState [11]string
	allowState[0] = "stop"
	allowState[1] = "clear"
	allowState[2] = "managed"
	allowState[3] = "ventilation"
	allowState[4] = "unfaithful"
	allowState[5] = "low_liquidity"
	allowState[6] = "lack_listed"
	allowState[7] = "overheated"
	allowState[8] = "caution"
	allowState[9] = "warning"
	allowState[10] = "risk"
	AllowState = allowState
}

type ViewPriceParms struct {
	Limit  int
	Offset int
	Sort   string
	Desc   bool
	Market []string
	State  map[string]bool
	Search string
}

func (obj *ViewPriceParms) GetDesc() string {
	if obj.Desc {
		return "desc"
	} else {
		return "asc"
	}
}

func (obj *ViewPriceParms) SetPageRows(page string, rows string) {

	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}

	limit, err := strconv.Atoi(rows)
	if err != nil {
		limit = Rows
	}
	offset := (p - 1) * limit

	obj.Limit = limit
	obj.Offset = offset
}

func (obj *ViewPriceParms) SetSortDesc(sort string, in_desc string) {

	desc, err := strconv.ParseBool(in_desc)
	if err != nil {
		desc = true
	}
	obj.Desc = desc
	obj.Sort = AllowSort[5]

	for _, s := range AllowSort {
		if s == sort {
			obj.Sort = sort
			break
		}
	}

}

func (obj *ViewPriceParms) SetEtc(market string, search string) {

	m_arr := make(map[string]bool)
	market_str := strings.TrimSpace(market)
	str := strings.Split(market_str, ",")
	for _, i := range AllowMarket {
		is := false
		for _, j := range str {
			if i == j {
				is = true
			}
		}
		m_arr[i] = is
	}
	keys := []string{}
	for key, value := range m_arr {
		if value {
			keys = append(keys, key)
		}
	}

	obj.Market = keys
	obj.Search = search
}

func (obj *ViewPriceParms) SetState(state string) {

	m_arr := make(map[string]bool)
	state_str := strings.TrimSpace(state)
	str := strings.Split(state_str, ",")

	for _, i := range AllowState {
		for _, j := range str {
			str_arr := strings.Split(j, "::")
			if len(str_arr) == 2 {
				if i == str_arr[0] {
					v, err := strconv.ParseBool(str_arr[1])
					if err == nil {
						m_arr[str_arr[0]] = v
					}
				}
			}

		}
	}
	obj.State = m_arr
}

type ViewPrice struct {
	Full_count int `json:"full_count"`

	Code       string          `json:"code"`
	Name       string          `json:"name"`
	Market     string          `json:"market"`
	High_date  sql.NullString  `json:"high_date"`
	High_price sql.NullFloat64 `json:"high_price"`

	Last_close_price sql.NullFloat64 `json:"last_close_price"`
	Contrast_price   sql.NullFloat64 `json:"contrast_price"`
	Fmt_high_date    sql.NullString  `json:"fmt_high_date"`
	Fmt_high_price   sql.NullString  `json:"fmt_high_price"`
	Fmt_last_date    sql.NullString  `json:"fmt_last_date"`

	Fmt_last_close_price    sql.NullString  `json:"fmt_last_close_price"`
	Fmt_contrast_price      sql.NullString  `json:"fmt_contrast_price"`
	Fluctuation_rate        sql.NullFloat64 `json:"fluctuation_rate"`
	Day_count               sql.NullInt64   `json:"day_count"`
	Updated_date_high_point string          `json:"updated_date_high_point"`

	Naver_link  string `json:"naver_link"`
	Stop        bool   `json:"stop"`
	Clear       bool   `json:"clear"`
	Managed     bool   `json:"managed"`
	Ventilation bool   `json:"ventilation"`

	Unfaithful    bool `json:"unfaithful"`
	Low_liquidity bool `json:"low_liquidity"`
	Lack_listed   bool `json:"lack_listed"`
	Overheated    bool `json:"overheated"`
	Caution       bool `json:"caution"`

	Warning                    bool   `json:"warning"`
	Risk                       bool   `json:"risk"`
	Updated_date_company_state string `json:"updated_date_company_state"`
}

type ViewMarket struct {
	ShortCode        string `json:"short_code"`
	High_date        string `json:"high_date"`
	High_price       string `json:"high_price"`
	Last_date        string `json:"last_date"`
	Last_close_price string `json:"last_close_price"`

	Contrast_price          string `json:"contrast_price"`
	Fluctuation_rate        string `json:"fluctuation_rate"`
	Day_count               int    `json:"day_count"`
	High_point_updated_date string `json:"high_point_updated_date"`
	Naver_link              string `json:"naver_link"`
}
