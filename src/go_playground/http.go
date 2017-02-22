package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

func SimpleHTTPServer() {
	port := ":8888"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)
		body, _ := ioutil.ReadAll(r.Body)
		log.Print(string(body))
		w.Write([]byte(`{"err_code":0, "err_msg":"OK"}`))
	})
	log.Print("listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	SimpleHTTPServer()
}
