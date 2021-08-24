package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/getTimes", getTimes)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func getTimes(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.Form["uid"]
	sid := r.Form["sid"]
	fmt.Println(uid, sid)
	fmt.Fprintf(w, "hello")
}
