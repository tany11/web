package database

import (
	"fmt"
	"log"
	"os"

	"back2/internal/domain/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func InitDatabase() (*xorm.Engine, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_USER_PASSWORD")
	dbHost := os.Getenv("DB_ADDRESS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbName)
	fmt.Printf("Connecting to database with DSN: %s\n", dsn)

	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create xorm engine: %w", err)
	}

	if err = engine.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	engine.ShowSQL(true)
	engine.SetMaxOpenConns(2)

	return engine, nil
}

func SyncDatabase(engine *xorm.Engine) error {
	tables := []interface{}{
		new(entity.Group),
		new(entity.Cast),
		new(entity.Customer),
		new(entity.CustomerStatus),
		new(entity.Orders),
		new(entity.Store),
		new(entity.Staff),
		new(entity.Media),
	}

	for _, table := range tables {
		if err := engine.Sync2(table); err != nil {
			return fmt.Errorf("failed to sync table %T: %w", table, err)
		}
	}

	log.Println("Database synchronization completed successfully")
	return nil
}
