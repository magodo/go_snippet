package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/udb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//_, _ = db.Exec(`INSERT INTO t_db_backup VALUES
	//(
	//	1, "0.0.0.0", 1, "0.0.0.0", "foo", 123, 0, 0, 0,0,0,0
	//)`)

	//fields := [...]string{"host_ip", "create_time", "backup_size"}
	//err = basic.UpdateMgrDB(db, "t_db_backup", 1, fields[:],
	//	"1.1.1.1", time.Now().Unix(), 123)
	//if err != nil {
	//	log.Fatal(err)
	//}
	rows, err := db.Query("show slave status")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	showResults := make(map[string]string)
	colValues := make([]interface{}, len(cols))
	for i := 0; i < len(colValues); i++ {
		colValues[i] = new(sql.RawBytes)
	}

	rows.Next()
	err = rows.Scan(colValues...)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(cols); i++ {
		showResults[cols[i]] = string(*colValues[i].(*sql.RawBytes))
	}

	spew.Dump(showResults)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(showResults["Slave_IO_Running"] == "Connecting")

}
