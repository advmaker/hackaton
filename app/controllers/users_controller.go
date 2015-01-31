package controllers

import (
	"github.com/jinzhu/gorm"
	"log"
)

type UsersController struct{ Controller }

func (controller *UsersController) New(db *gorm.DB, l *log.Logger) string {
	return "some"
}
