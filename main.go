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
	bsid := r.Form["sid"]
	var sid string

	switch bsid[0] {
	case "6006", "6007":
		sid = "6005"
	case "6009", "6010":
		sid = "6008"
	case "6012", "6013", "6014", "6015", "6016", "6017":
		sid = "6011"
	case "6019", "6020":
		sid = "6018"
	default:
		sid = bsid[0]
	}

	fmt.Println(uid[0], bsid[0], sid)

	cmd := fmt.Sprintf("LC_ALL=C fgrep %s %s/%s/* | wc -l", uid[0], logDir, sid)
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		log.Panic(err)
	}

	//cmdTotal := fmt.Sprintf("wc -l %s/*/* | grep total | awk -F ' ' '{print $1}'", logDir)
	/*
		outTotal, err := exec.Command("/bin/bash", "-c", cmdTotal).Output()
		if err != nil {
			log.Panic(err)
		}
	*/

	fmt.Fprint(w, string(out))
}
