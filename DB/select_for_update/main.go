package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/foobar")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.BeginTx(context.TODO(), nil)

	var sqlErr error
	defer func() {
		if sqlErr != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var name string
	if err := tx.QueryRow(`select name from a where id = 0 for update`).Scan(&name); err != nil {
		sqlErr = err
		return
	}
}
