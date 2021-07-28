package handler

import (
	"net/http"

	"github.com/cheolgyu/stock-read-pub-api/src/config"
	"github.com/julienschmidt/httprouter"
)

type CheckHandler struct {
	//AUsecase domain.ArticleUsecase
}

func NewCheckHandler(h httprouter.Handle) {

	hadler := CheckHandler{}
	hadler.is_market_code(h)

}

// 마켓 상세도 나와야 대니깐. 추가한 함수였네
// 마켓과 종목이 별도의 테이블에 되있었지
func ChkMarketCode(code string) (bool, int) {
	is_market := false
	idx := -1
	for i := range config.MarketList {
		if config.MarketList[i].Code == code {
			idx = i
			is_market = true
			break
		}
	}
	return is_market, idx
}

func (obj *CheckHandler) is_market_code(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		req_code := ps.ByName("code")
		ok, _ := ChkMarketCode(req_code)

		if ok {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
