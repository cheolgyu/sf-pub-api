package model

import (
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

func ChkMarket(market string) bool {

	market_str := strings.TrimSpace(market)
	chk := false
	for i := range AllowMarket {
		if AllowMarket[i] == market_str {
			chk = true
			break
		}
	}

	return chk
}

func ChkDate(str string) bool {
	chk := false
	res, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		return false
	}

	if res > 0 && res < 30000000 {
		chk = true
	}

	return chk
}
