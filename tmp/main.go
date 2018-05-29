package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {

	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/udb")
	if err != nil {
		log.Fatal(err)
	}

	outs, err := Vquery("show variables like 'log_timestamps'", db)
	if err != nil {
		log.Fatal(err)
	}
	if outs == nil {
		fmt.Println("No match")
	}
}

func Vquery(query string, db *sql.DB) (results []map[string]string, err error) {
	rows, err := db.Query(query)
	if err != nil {
		err = errors.Wrapf(err, "query %s failed", query)
		return
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		err = errors.Wrap(err, "get column failed")
		return
	}

	for rows.Next() {
		colValues := make([]interface{}, len(cols))
		for i := 0; i < len(colValues); i++ {
			colValues[i] = new(sql.RawBytes)
		}

		err = rows.Scan(colValues...)
		if err != nil {
			err = errors.Wrap(err, "scan slave status failed")
			return
		}

		result := make(map[string]string, len(cols))
		for i := 0; i < len(cols); i++ {
			result[cols[i]] = string(*colValues[i].(*sql.RawBytes))
		}
		results = append(results, result)
	}

	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "something failed during slave status rows operation")
		return
	}
	return
}
