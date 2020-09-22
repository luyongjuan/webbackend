package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"luyongjuan/webbackend/repository"
	"math/rand"
)

func initDB(host, user, pw, dbPort string) *gorm.DB {

	db := repository.InitDB(host, user, pw, dbPort)
	if db != nil {
		//init db interface
	}

	return db
}

func main() {

	/*dbHost := os.Getenv("POSTGRES_URL")
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

	http.ListenAndServe(":8080", nil)*/

	var data []int
	for i := 10; i < 731; i++{
		data = append(data, rand.Intn(5)+rand.Intn(10))
	}

	fmt.Println(data)


}
