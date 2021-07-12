package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cheolgyu/stock-read-pub-api/src/service"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

var UPDATED string
var DATA []byte
var frontend_url string
var port string

var MarketList = []string{"KOSPI", "KOSDAQ", "FUT", "KPI200"}
var MarketListName = []string{"코스피", "코스닥", "선물", "코스피200"}

func Exec(isDebug bool) {
	frontend_url = os.Getenv("FRONTEND_URL")
	port = ":" + os.Getenv("PORT")
	log.Println("frontend_url", frontend_url)
	server()
}

type Middleware struct {
	next    http.Handler
	message string
}

func NewMiddleware(next http.Handler, message string) *Middleware {
	return &Middleware{next: next, message: message}
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	req_id := uuid.New().String()
	r.Header.Add("req_id", req_id)
	log.Printf("[Middleware] <%s> %s %s %s\n", req_id, r.RemoteAddr, r.Method, r.URL)
	m.next.ServeHTTP(w, r)
	log.Printf("[Middleware] <%s> %s \n", req_id, w.Header())
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req_id := r.Header.Get("req_id")
	log.Printf("<%s> \n ", req_id)
	msg := fmt.Sprintf("Welcome! %s\n", time.Now().String())
	fmt.Fprint(w, msg)
}

type ViewPriceResult struct {
	Info   []map[string]interface{} `json:"info"`
	Price  []map[string]interface{} `json:"price"`
	Market []map[string]interface{} `json:"market"`
}

func HandlerPriceBound(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	req_code := ps.ByName("code")
	tbnm := "hist.bound_stock"
	market, _ := ChkMarketCode(req_code)
	if market {
		tbnm = "hist.bound_market"
	}

	list := service.GetHistPriceBound(req_id, r, tbnm, req_code)
	json.NewEncoder(w).Encode(list)

}

func HandlerMonthlyPeek(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)
	list := service.GetMonthlyPeek(req_id, r)
	json.NewEncoder(w).Encode(list)
}

func HandlerDayTrading(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)
	list := service.GetDayTrading(req_id, r)
	json.NewEncoder(w).Encode(list)
}

func HandlerInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)
	res := ViewPriceResult{}
	info := service.GetInfo(req_id)
	res.Info = info
	json.NewEncoder(w).Encode(res)
}

func HandlerViewPrice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	list := service.GetViewPrice(req_id, r)
	info := service.GetInfo(req_id)
	res := ViewPriceResult{}
	res.Info = info
	res.Price = list
	json.NewEncoder(w).Encode(res)

}
func HandlerViewMarket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	info := service.GetInfo(req_id)
	market_list := service.GetMarket(req_id, r)
	//log.Println(market_list)
	res := ViewPriceResult{}
	res.Info = info
	res.Market = market_list
	json.NewEncoder(w).Encode(res)

}

func ChkMarketCode(code string) (bool, int) {
	is_market := false
	idx := -1
	for i := range MarketList {
		if MarketList[i] == code {
			idx = i
			is_market = true
			break
		}
	}
	return is_market, idx
}

func HandlerDetailChart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	req_code := ps.ByName("code")
	q := r.URL.Query()
	page := q.Get("page")

	tbnm := "hist.price_stock"
	market, _ := ChkMarketCode(req_code)
	if market {
		tbnm = "hist.price_market"
	}

	pnum, err := strconv.Atoi(page)
	if err != nil {
		pnum = 1
	}

	list := service.GetDetailChart(req_id, tbnm, req_code, pnum)

	json.NewEncoder(w).Encode(list)

}

func HandlerDetailChartLine(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	req_code := ps.ByName("code")

	list := service.GetDetailChartLine(req_id, req_code)

	json.NewEncoder(w).Encode(list)

}

func HandlerDetailCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	req_code := ps.ByName("code")
	var list string
	market, i := ChkMarketCode(req_code)

	if market {
		list = `{ "c": { "code": "` + MarketList[i] + `","name":"` + MarketListName[i] + `"}}`
	} else {
		list = service.GetDetailCompany(req_id, req_code)
	}

	json.NewEncoder(w).Encode(list)

}

func setCors(w *http.ResponseWriter) {
	header := (*w).Header()

	header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header.Set("Access-Control-Allow-Origin", frontend_url)
	header.Set("Access-Control-Allow-Credentials", "true")
	header.Set("Content-Type", "application/json")
}

func server() {
	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {

			header := w.Header()
			header.Set("Access-Control-Allow-Origin", frontend_url)
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Credentials", "true")
		}
		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/", Index)
	router.GET("/info", HandlerInfo)
	router.GET("/price", HandlerViewPrice)
	router.GET("/price/bound/:code", HandlerPriceBound)
	router.GET("/market", HandlerViewMarket)
	router.GET("/detail/chartline/:code", HandlerDetailChartLine)
	router.GET("/detail/chart/:code", HandlerDetailChart)
	router.GET("/detail/company/:code", HandlerDetailCompany)
	router.GET("/day_trading", HandlerDayTrading)
	router.GET("/monthly_peek", HandlerMonthlyPeek)
	m := NewMiddleware(router, "I'm a middleware")
	log.Fatal(http.ListenAndServe(port, m))
}
