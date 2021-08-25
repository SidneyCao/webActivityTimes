package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var logDir string = "/data/webActivityLottery/"

func main() {
	http.HandleFunc("/getTimes", getTimes)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func getTimes(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.Form["uid"]
	sid := r.Form["sid"]
	fmt.Println(uid[0], sid[0])

	cmd := fmt.Sprintf("LC_ALL=C fgrep %s %s/%s/* | wc -l", uid[0], logDir, sid[0])
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()

	if err != nil {
		log.Panic(err)
	}

	fmt.Fprint(w, string(out))
}
