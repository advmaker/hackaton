
package main

import (
  "database/sql"

  "github.com/jinzhu/gorm"

  "github.com/advmaker/hackaton/config"
  "github.com/advmaker/hackaton/app/models"
)

var db gorm.DB = config.DB()

func Up_20150114090334(txn *sql.Tx) {
  db.CreateTable(&models.Progress{})
}

func Down_20150114090334(txn *sql.Tx) {
  db.DropTable(&models.Progress{})
}
