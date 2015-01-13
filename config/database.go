package config

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

func DB() gorm.DB {
  db, err := gorm.Open("postgres", "user=achievements_user dbname=achievements_dev_db password=achievement_unlock! sslmode=disable")
  if err != nil { log.Fatal(err) }

  db.DB().Ping()
  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)

  return db
}

func CloseDB(db gorm.DB) {

}
