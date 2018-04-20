package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/magodo/go_snippet/DB/mysql/basic"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/udb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, _ = db.Exec(`INSERT INTO t_db_backup VALUES
	(
		1, "0.0.0.0", 1, "0.0.0.0", "foo", 123, 0, 0, 0,0,0,0
	)`)

	fields := [...]string{"host_ip", "create_time", "backup_size"}
	err = basic.UpdateMgrDB(db, "t_db_backup", 1, fields[:],
		"1.1.1.1", time.Now().Unix(), 123)
	if err != nil {
		log.Fatal(err)
	}
}
