package main

import (
	"log"

	"github.com/cheolgyu/pubapi/src/server"
)

/*

gin -i --appPort 5001  --port 5000  run -- -prod main.go
gin -i --appPort 5001  --port 5000  run  main.go
*/

var isDebug bool = true

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

// go run data-server/main.go
func main() {
	log.Println("데이터 서버 시작")
	server.Exec(isDebug)
}
