package basic

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

func TransactionalInsert(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	stmt, err := tx.Prepare("insert into `zoo`(kind, name, age) values(?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("小凶许", "diudiu", 3)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec("小脑斧", "dd", 3)
	if err != nil {
		log.Fatalln(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

func QueryZoo(db *sql.DB) {
	stmt, err := db.Prepare("select kind from `zoo`")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var kind string
	for rows.Next() {
		err := rows.Scan(&kind)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(kind)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func QueryNonexistance(db *sql.DB) {
	stmt, err := db.Prepare("select kind from `zoo` where name = 'foo'")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var kind string
	err = stmt.QueryRow().Scan(&kind)
	if err != nil {
		// special handle "no row" error
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
	}
}

func QueryNull(db *sql.DB) {
	stmt, err := db.Prepare("select * from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		name sql.NullString
		id   int
	)
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		if name.Valid {
			fmt.Printf("%5d %10s\n", id, name)
		} else {
			fmt.Printf("name of %d is empty\n", id)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateMgrDB(db *sql.DB, table string, id int64, fields []string, values ...interface{}) (err error) {

	if len(fields) != len(values) {
		err = errors.New("insane argument")
		return
	}

	stmt := fmt.Sprintf("update %s set %s where id = %d", table, strings.Join(fields, "=?,")+"=?", id)

	if strings.Count(stmt, "?") != len(values) {
		err = errors.New(fmt.Sprintf("malformed sql statement %s, should not contain '?' as non-place holder", stmt))
		return
	}

	fullStmt := stmt
	for i := 0; i < len(values); i++ {
		fullStmt = strings.Replace(fullStmt, "?", fmt.Sprint(values[i]), 1)
	}

	fmt.Println(fullStmt)
	_, err = db.Exec(stmt, values...)
	if err != nil {
		err = errors.Wrapf(err, "exec SQL: %s failed", fullStmt)
	}
	return
}

func ShowSQLStmt(stmt string, args ...interface{}) (fullStmt string, err error) {
	if strings.Count(stmt, "?") != len(args) {
		err = errors.New(fmt.Sprintf("malformed sql statement %s, should not contain '?' as non-place holder", stmt))
		return
	}

	fullStmt = stmt
	for i := 0; i < len(args); i++ {
		fullStmt = strings.Replace(fullStmt, "?", fmt.Sprintf("%q", args[i]), 1)
	}

	return
}
