package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(179.17.0.2:3306)/mysql?timeout=3s")
	if err != nil {
		log.Fatal(err)
	}
	for {
		for n := 0; n < 10; n++ {
			ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
			_, err = db.ExecContext(ctx, "select 1")
			if err != nil {
				err = errors.Wrap(err, "failed to query")
				fmt.Println(err)
			} else {
				fmt.Println("query OK")
			}
			time.Sleep(time.Second)
		}
		fmt.Println("loop finish")

		//ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		//err = db.PingContext(ctx)
		//if err != nil {
		//	err = errors.Wrap(err, "ping failed")
		//	fmt.Println(err)
		//} else {
		//	fmt.Println("ping OK")
		//}
		//time.Sleep(time.Second)
	}

	//name := ""
	//err = db.QueryRow("SELECT name FROM zoo where age = 12").Scan(&name)
	//if err != nil {
	//	if err != sql.ErrNoRows {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("no row...")
	//}
	//_, err = db.Exec("INSERT INTO bar(val) values(?)", uint64(math.MaxUint64))
	//if err != nil {
	//	log.Fatal(err)
	//}

	//------------------------------------------
	//row, err := db.Query("SELECT name, age from zoo")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//defer row.Close()
	//for row.Next() {
	//	var name string
	//	var age int
	//	err := row.Scan(&name, &age)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	//if !name.Valid {
	//	//	fmt.Println("name is nil when agen = ", age)
	//	//} else {
	//	//	fmt.Println("name = ", name)
	//	//}
	//	fmt.Println("name = ", name)
	//}

	//-----------------------------------------
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

	//-----------------------------------------
	//rows, err := db.Query("show slave status")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()

	//cols, err := rows.Columns()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//showResults := make(map[string]string)
	//colValues := make([]interface{}, len(cols))
	//for i := 0; i < len(colValues); i++ {
	//	colValues[i] = new(sql.RawBytes)
	//}

	//rows.Next()
	//err = rows.Scan(colValues...)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//for i := 0; i < len(cols); i++ {
	//	showResults[cols[i]] = string(*colValues[i].(*sql.RawBytes))
	//}

	//spew.Dump(showResults)
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(showResults["Slave_IO_Running"] == "Connecting")
	//result, err := mylib.Vquery(db, "show slave status")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//spew.Dump(result)

}
