package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"strconv"
	"time"
)


func InitDB(host, user, pw, dbPort string)*gorm.DB{


	port, err := strconv.Atoi(dbPort)
	if err != nil{
		panic(err)
		return nil
	}


	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pw, "operation_db")

	fmt.Println(dbUrl)

	var db *gorm.DB
	for {
			db, err = gorm.Open("postgres", dbUrl)
			if err !=nil{
				log.Print(err.Error())
				time.Sleep(3*time.Second)
			}else {
				break
			}
	}

	return db
}
