package initializers

import (
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

func init() {
  db, err := gorm.Open("postgres", "user=achievements_user dbname=achievements_dev_db password=achievement_unlock! sslmode=disable")

  if err == nil {}

  db.DB().Ping()
  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)
}

