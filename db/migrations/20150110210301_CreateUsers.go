package main

import (
	"database/sql"
  "github.com/jinzhu/gorm"
  "github.com/advmaker/hackaton/achievement/models"
  "github.com/advmaker/hackaton/config"
)

var db gorm.DB = config.DB()

func Up_20150110210301(txn *sql.Tx) {
  db.CreateTable(&models.User{})
}

func Down_20150110210301(txn *sql.Tx) {
  db.DropTable(&models.User{})
}
