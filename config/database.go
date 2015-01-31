package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func DB() gorm.DB {
	db, err := gorm.Open(os.Getenv("DB_ADAPTER"), os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}

func CloseDB(db gorm.DB) {

}
