package controllers

import (
	"github.com/advmaker/hackaton/models"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/martini"
)

type Controller struct{}

func (*Controller) auth(params martini.Params, db *gorm.DB) {
	auth.BasicFunc(func(username, password string) bool {
		db.Where(&User{email: params["email"]}).First(&user)
		if user != nil {
			auth.SecureCompare(user.Password, params["password"])
		}
	})
}
