package controllers

import (
	"github.com/go-martini/martini"
)

type UsersController struct{ Controller }

func (controller *UsersController) New(l martini.Logger) string {

	return "some"
}
