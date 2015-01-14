package main

import (
	"database/sql"

  "github.com/jinzhu/gorm"

  "github.com/advmaker/hackaton/config"
  "github.com/advmaker/hackaton/app/models"
)

var db gorm.DB = config.DB()

func Up_20150110210301(txn *sql.Tx) {
  db.CreateTable(&models.User{})
}

func Down_20150110210301(txn *sql.Tx) {
  db.DropTable(&models.User{})
}
