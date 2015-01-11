package main

import (
  "log"
	"database/sql"
  "github.com/jinzhu/gorm"
  "github.com/advmaker/hackaton/achievement/models"
  _ "github.com/advmaker/hackaton/config/initializers"
)

func Up_20150110210301(txn *sql.Tx) {
  db, err := gorm.Open("postgres", "user=achievements_user dbname=achievements_dev_db password=achievement_unlock! sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
  db.CreateTable(&models.User{})
}

func Down_20150110210301(txn *sql.Tx) {
  db, err := gorm.Open("postgres", "user=achievements_user dbname=achievements_dev_db password=achievement_unlock! sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
  db.DropTable(&models.User{})
}
