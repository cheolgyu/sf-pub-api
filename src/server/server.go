package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cheolgyu/stock-read-pub-api/src/db"
	"github.com/cheolgyu/stock-read-pub-api/src/svc/company"
	"github.com/cheolgyu/stock-read-pub-api/src/svc/info"
	"github.com/cheolgyu/stock-read-pub-api/src/svc/meta"
	"github.com/cheolgyu/stock-read-pub-api/src/svc/price"
	"github.com/cheolgyu/stock-read-pub-api/src/svc/project"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

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
	setCors(&w)
	m.next.ServeHTTP(w, r)

	//log.Printf("[Middleware] <%s> %s \n", req_id, w.Header())
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req_id := r.Header.Get("req_id")
	log.Printf("<%s> \n ", req_id)
	msg := fmt.Sprintf("Welcome! %s\n", time.Now().String())
	fmt.Fprint(w, msg)
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
	m := NewMiddleware(router, "I'm a middleware")

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {

			header := w.Header()
			header.Set("Access-Control-Allow-Origin", frontend_url)
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Credentials", "true")
		}
		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/", Index)
	timeoutContext := time.Duration(2) * time.Second

	db_conn := db.Conn()

	info_repo := info.NewRepository(db_conn)
	info_usecase := info.NewUsecase(info_repo, timeoutContext)
	info.NewHandler(router, info_usecase)

	meta_repo := meta.NewRepository(db_conn)
	meta_usecase := meta.NewUsecase(meta_repo, timeoutContext)
	meta.NewHandler(router, meta_usecase)

	cmp_repo := company.NewRepository(db_conn)
	cmp_usecase := company.NewUsecase(cmp_repo, meta_repo, timeoutContext)
	company.NewHandler(router, cmp_usecase)

	price_repo := price.NewRepository(db_conn)
	price_usecase := price.NewUsecase(price_repo, meta_repo, timeoutContext)
	price.NewHandler(router, price_usecase)

	project_repo := project.NewRepository(db_conn)
	project_usecase := project.NewUsecase(project_repo, meta_repo, timeoutContext)
	project.NewHandler(router, project_usecase)

	log.Fatal(http.ListenAndServe(port, m))
}
