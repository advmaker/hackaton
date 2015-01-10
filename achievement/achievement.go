package achievement

import (
	"github.com/go-martini/martini"
//	"github.com/advmaker/hood"
//	"fmt"
)

type Application struct {
	martini *martini.ClassicMartini
}

func (a *Application) Run() {
	a.martini = martini.Classic()
	a.martini.Get("/", func() string {
			return "some"
		})
	a.martini.Run()
}
