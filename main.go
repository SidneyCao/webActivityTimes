package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var logDir string = "/data/webActivity/"

func main() {
	http.HandleFunc("/getTimes", getTimes)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func getTimes(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.Form["uid"]
	sid := r.Form["sid"]
	fmt.Println(uid[0], sid[0])

	cmd := fmt.Sprintf("LC_ALL=C fgrep lottery %s/%s/* | LC_ALL=C fgrep '{\"jid\":\"109'| wc -l", logDir, sid[0])
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()

	if err != nil {
		log.Panic(err)
	}

	fmt.Fprint(w, string(out))
}
