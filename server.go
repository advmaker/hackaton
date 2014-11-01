package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/auth"
	"github.com/eaigner/hood"
)

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
		Directory: "templates",
	}))

	m.Use(martini.Static("public"))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "parts/index", "Anna")
	})

	m.Run()
}
