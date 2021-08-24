package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	logDir string = "/data/webActivity/"
)

func main() {
	http.HandleFunc("/getTimes", getTimes)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func getTimes(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.Form["uid"]
	sid := r.Form["sid"]
	fmt.Println(uid[0], sid[0])
	fmt.Fprintf(w, "hello")
}
