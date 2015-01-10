package achievement

import (
	"github.com/go-martini/martini"
)

type Application struct {
	martini *martini.ClassicMartini
}

func (a *Application) Run() {
	a.martini = martini.Classic()
	a.martini.Get("/", func() string {
			return "555"
		})
	a.martini.Run()
}
