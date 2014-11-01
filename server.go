package main

import (
//	"github.com/go-martini/martini"
//	"github.com/martini-contrib/render"
	"github.com/eaigner/hood"
	"log"
)

func main() {

	db, err := hood.Open("postgres", "host=127.0.0.1 port=5432 dbname=achievements_db user=root password=аа sslmode=verify-full")
	if err != nil {
//		log.Fatal(err)
	} else {
		some := db.QueryRow("SELECT (1+1) AS some")
		log.Fatal(some)
	}
}
