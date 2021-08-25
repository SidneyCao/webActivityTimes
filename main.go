package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

var (
	logDir string = "/data/webActivity/"
	DBhost        = flag.String("h", "", "host 默认为空")
	DBuser        = flag.String("u", "root", "user 默认为root")
	DBpwd         = flag.String("p", "", "password 默认为空")
)

func main() {
	//获取命令行参数
	flag.Parse()

	session, err := getDBsession()
	if err != nil {
		log.Panicf("创建数据库会话失败: %w", err)
	}

	http.HandleFunc("/getTimes", getTimes)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func getTimes(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.Form["uid"]
	sid := r.Form["sid"]
	fmt.Println(uid[0], sid[0])

}

func getDBsession() (*mgo.Session, error) {
	session, err := mgo.Dial(*DBhost)
	if err != nil {
		return nil, fmt.Errorf("无法连接数据库: %w", err)
	}
	session.SetMode(mgo.Monotonic, true)
	adminDB := session.DB("admin")
	err = adminDB.Login(*DBuser, *DBpwd)
	if err != nil {
		return nil, fmt.Errorf("登录失败: %w", err)
	}
	return session, nil
}
