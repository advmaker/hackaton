package main

import (
	//  "fmt"
	"github.com/advmaker/hackaton/app/controllers"
	"github.com/advmaker/hackaton/config"

	"github.com/go-martini/martini"
	_ "github.com/joho/godotenv/autoload"
	"github.com/martini-contrib/render"
)

func main() {
	app := martini.Classic()

	usersController := new(controllers.UsersController)


	renderOptions := render.Options{
		Layout: "layout",
	}

	app.Use(render.Renderer(renderOptions))

	app.Use(martini.Logger())

	app.Map(config.DB())

	app.Get("/users/new", usersController.New)

	app.Run()
}
