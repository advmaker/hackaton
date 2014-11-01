package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	//	"github.com/go-martini/martini"
	//	"github.com/martini-contrib/render"
	//  "github.com/martini-contrib/auth"
	//	"github.com/eaigner/hood"
	"log"
	"fmt"
)

func main() {
	//	db, err := sql.Open("postgres", "postgres://root:G4rr6Df7l@hotfile25.com")
	db, err := sql.Open("postgres", "host=hotfile25.com port=5432 dbname=achievements_db user=root password=G4rr6Df7l")
	//	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=achievements_db user=vagrant password=vagrant")
	if err != nil {
		log.Fatal(err)
	} else {
		var s string
		db.QueryRow("SELECT 1+1 as c").Scan(&s)
		fmt.Println(s)
	}
}
