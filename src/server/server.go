package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/cheolgyu/stock/backend/api/src/model"
	"github.com/cheolgyu/stock/backend/api/src/service"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

var UPDATED string
var DATA []byte
var frontend_url string
var port string

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
	fmt.Fprint(w, "Welcome!\n")
}

type ViewPriceResult struct {
	Info   map[string]string  `json:"info"`
	Price  []model.ViewPrice  `json:"price"`
	Market []model.ViewMarket `json:"market"`
}

func HandlerViewPrice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	list := service.GetViewPrice(req_id, r)
	info := service.GetInfo(req_id)
	market_list := service.GetMarket(req_id)

	res := ViewPriceResult{}
	res.Info = info
	res.Price = list
	res.Market = market_list
	json.NewEncoder(w).Encode(res)

}

func HandlerDetailChart(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	req_code := ps.ByName("code")
	q := r.URL.Query()
	page := q.Get("page")

	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}

	list := service.GetDetailChart(req_id, req_code, p)

	json.NewEncoder(w).Encode(list)

}

func HandlerDetailCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req_id := r.Header.Get("req_id")
	setCors(&w)

	req_code := ps.ByName("code")

	list := service.GetDetailCompany(req_id, req_code)

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
	router.GET("/price", HandlerViewPrice)
	router.GET("/detail/chart/:code", HandlerDetailChart)
	router.GET("/detail/company/:code", HandlerDetailCompany)
	m := NewMiddleware(router, "I'm a middleware")
	log.Fatal(http.ListenAndServe(port, m))
}
