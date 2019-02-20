package main

import (
	"errors"
	"github.com/jinzhu/gorm"
	"luyongjuan/webbackend/datapush"
	"luyongjuan/webbackend/handler"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"luyongjuan/webbackend/repository"

)

func initDB(host, user, pw, dbPort string) *gorm.DB {

	db := repository.InitDB(host, user, pw, dbPort)
	if db != nil {
		//init db interface
	}

	return db
}

func main() {

	dbHost := os.Getenv("POSTGRES_URL")
	if dbHost == "" {
		dbHost = "172.17.0.2"
	}

	username := os.Getenv("POSTGRES_USER")
	if username == "" {
		username = "postgres"
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "root"
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	gdb := initDB(dbHost, username, password, dbPort)
	if gdb == nil {
		panic(errors.New("can not connect to postgres"))
		return
	}

	// go kit init
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	httpLogger := log.With(logger, "component", "http")

	//websocket
	datapush.DataPusherInit()
	//
	test := handler.NewTestService()


	//http request init
	mux := http.NewServeMux()
	//consumer insight
	mux.Handle("/tt/", handler.MakeHandler(test, httpLogger))
	http.Handle("/", mux)

	http.ListenAndServe(":8080", nil)
}
