package main

import (
	"fmt"
	"time"
	"net/http"
	"log"
	"regexp"
	"os"
)


var logger *log.Logger
var URL_PATH_FORMAT string = "/static/data_for_%s.tar.gz";

func initLogger() {
	log_file, _ := os.OpenFile("./info.log", os.O_RDWR | os.O_CREATE, 777)
	logger = log.New(log_file, "", log.Ldate | log.Llongfile | log.Ltime)
}
func ExchangeUser() {
	fmt.Println("test")
}

func LoadData() {
	ExchangeUser()
}
func Tick() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("loading data, %v", now)
		LoadData()
	}
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
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	if ok, _ := regexp.MatchString("/static/", r.URL.String()); ok {
		StaticServerHandler(w, r)
		return
	}
	http.Error(w, "", http.StatusInternalServerError)
}


func StaticServerHandler(w http.ResponseWriter, r *http.Request) {
	logger.Print(r.URL.Path)
	if (r.URL.RawQuery != "") {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	src := "jd"
	required_url_path := fmt.Sprintf(URL_PATH_FORMAT, src)
	if (r.URL.Path != required_url_path) {
		err_msg := fmt.Sprintf("invalid path: %s", r.URL.Path)
		logger.Print(err_msg)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "", http.StatusInternalServerError)
	}
	wd = fmt.Sprintf("%s/data", wd)
	http.StripPrefix("/static/",
		http.FileServer(http.Dir(wd))).ServeHTTP(w, r)
}
