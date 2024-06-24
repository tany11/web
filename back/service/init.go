package service

import (
	"back/models"
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
	fmt.Printf("Connecting to database with DSN: %s\n", DsName)
	var err error
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = DbEngine.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	if err := DbEngine.Sync2(new(models.GroupLogin)); err != nil {
		log.Fatalf("Failed to sync GroupLogin: %v", err)
	}
	if err := DbEngine.Sync2(new(models.CastLogin)); err != nil {
		log.Fatalf("Failed to sync CastLogin: %v", err)
	}
	if err := DbEngine.Sync2(new(models.Customer)); err != nil {
		log.Fatalf("Failed to sync Customer: %v", err)
	}
	if err := DbEngine.Sync2(new(models.Order)); err != nil {
		log.Fatalf("Failed to sync Customer: %v", err)
	}
	if err := DbEngine.Sync2(new(models.Store)); err != nil {
		log.Fatalf("Failed to sync Customer: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed to synchronize database table: %v", err)
	}
	fmt.Println("init data base ok")
}
