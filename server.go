package main

import (
  //  "fmt"
  "github.com/go-martini/martini"
  "github.com/advmaker/hackaton/config"
  "github.com/advmaker/hackaton/achievement/controllers"
)

func main() {
  app := martini.Classic()

  app.Map(config.DB())

  app.Get("/", controllers.ApplicationIndex)

  app.Run()
}
