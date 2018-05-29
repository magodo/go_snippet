package mylib

import (
	"database/sql"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestShowSlaveStatus(t *testing.T) {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/udb")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	results, err := Vquery(db, "show slave status")
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(results)
}
