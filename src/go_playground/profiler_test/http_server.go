package main

import (
	"time"
	"net/http"
	"log"
	"os"
	"strconv"
	_ "net/http/pprof"
)


var logger *log.Logger
var URL_PATH_FORMAT string = "/static/data_for_%s.tar.gz";

func initLogger() {
	log_file, _ := os.OpenFile("./info.log", os.O_RDWR | os.O_CREATE, 777)
	logger = log.New(log_file, "", log.Ldate | log.Llongfile | log.Ltime)
}

func Calc(n1 int64, n2 int64) int64 {
	var s int64 = 0
	var i int64 = 0
	for i = 0; i < n1; i++ {
		s += 1
	}
	i = 0
	for i = 0; i < n2; i++ {
		s += 1
	}
	return s
}

func main() {
	initLogger()
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", DefaultHandler)

	server := http.Server{
		Addr:        ":8080",
		Handler:     serveMux,
		ReadTimeout: 5 * time.Second,
	}

	logger.Print("serveing...")
	go func() {
		log.Println(http.ListenAndServe("localhost:8081", nil))
	}()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	var queries = r.URL.Query()
	n1, err1 := strconv.ParseInt(queries.Get("n1"), 10, 64)
	n2, err2 := strconv.ParseInt(queries.Get("n2"), 10, 64)
	if err1 != nil || err2 != nil {
		w.Write([]byte("error"))
		return
	}
	sum := Calc(n1, n2)
	sum = Calc(n1, n2)
	sum = Calc(n1, n2)
	sumStr := strconv.FormatInt(sum, 10)
	rsp := "result:" + sumStr
	w.Write([]byte(rsp))
}
