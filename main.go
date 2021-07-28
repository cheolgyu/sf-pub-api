package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/cheolgyu/stock-read-pub-api/src/server"
	"github.com/joho/godotenv"
)

/*

gin -i --appPort 5001  --port 5000  run -- -prod main.go
gin -i --appPort 5001  --port 5000  run  main.go
*/

var isDebug bool = true

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Panic("Error loading .env file")
	}
}

// go run data-server/main.go
func main() {
	logPath := "logs/api/development.log"

	openLogFile(logPath)

	log.Println("데이터 서버 시작")
	server.Exec(isDebug)
}

func openLogFile(logfile string) {

	if logfile != "" {
		t := time.Now()
		dirname := t.Format("2006-01-02")
		filename := t.Format("2006-01-02_15_04_05_000Z")
		os.MkdirAll("./logs/api/"+dirname, 0755)
		lf, err := os.OpenFile("./logs/api/"+dirname+"/"+filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		multi := io.MultiWriter(lf, os.Stdout)
		log.SetOutput(multi)
	}
}
