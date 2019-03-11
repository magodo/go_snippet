package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"user='%s' password='%s' host=%s port=%s dbname=postgres sslmode=disable connect_timeout=5",
			"postgres",
			"",
			"127.0.0.1",
			"5432",
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		//err = db.Ping()
		_, err = db.Exec(";")
		if err != nil {
			spew.Dump(err)
		} else {
			log.Println("OK")
		}
		time.Sleep(3 * time.Second)
	}
}
