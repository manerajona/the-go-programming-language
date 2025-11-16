package main

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	for i := 0; i < 5; i++ {
		Run()
	}
}

var db *sql.DB
var o sync.Once

func Run() {
	o.Do(func() {
		log.Println("opening connection to database")
		var err error
		db, err = sql.Open("sqlite3", "./mydb.db")
		if err != nil {
			log.Fatal(err)
		}
	})
}
