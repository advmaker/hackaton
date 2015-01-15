
package main

import (
  "database/sql"

  "github.com/jinzhu/gorm"

  "github.com/advmaker/hackaton/config"
  "github.com/advmaker/hackaton/app/models"
)

var db gorm.DB = config.DB()

func Up_20150114085245(txn *sql.Tx) {
  db.CreateTable(&models.Achievement{})
}

func Down_20150114085245(txn *sql.Tx) {
  db.DropTable(&models.Achievement{})
}
