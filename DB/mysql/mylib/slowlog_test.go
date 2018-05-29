package mylib

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestSlowQuery(t *testing.T) {

	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/udb")
	if err != nil {
		t.Fatal(err)
	}

	out, err := ReadSlowQueryFromTable(db, "2006-01-02T15:04:05", time.Now().Add(-24*time.Hour), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}
