package service

import (
	"back/model"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_USER_PASSWORD")
	DbHost := os.Getenv("DB_ADDRESS")
	DbName := os.Getenv("DB_NAME")
	DsName := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8", DbUser, DbPassword, DbHost, DbName)
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	err = DbEngine.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	err = DbEngine.Sync2(new(model.Book))
	if err != nil {
		log.Fatalf("Failed to synchronize database table: %v", err)
	}
	fmt.Println("init data base ok")
}
