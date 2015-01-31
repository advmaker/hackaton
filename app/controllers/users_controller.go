package controllers

import (
	"github.com/jinzhu/gorm"
	"log"

	"github.com/martini-contrib/render"
)

type UsersController struct{ Controller }

func (controller *UsersController) New(db *gorm.DB, l *log.Logger, r render.Render) {
	r.HTML(200, "hello", "oggggg")
}
