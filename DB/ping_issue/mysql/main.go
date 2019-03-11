package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	//_ "gitlab.ucloudadmin.com/udb/uframework/utils/mysql/mysql" // old version
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/foobar")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		err = db.Ping()
		//_, err = db.Ping("select 1")
		if err != nil {
			spew.Dump(err)
		} else {
			log.Println("OK")
		}
		time.Sleep(3 * time.Second)
	}
}
