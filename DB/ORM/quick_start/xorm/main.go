package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:123@tcp(127.0.0.1:3306)/orm_test")
	if err != nil {
		log.Fatal(err)
	}

	// set logger
	engine.ShowSQL(true)

	//var buf bytes.Buffer

	f, _ := os.Open("foo.log")
	defer f.Close()
	logger := xorm.NewSimpleLogger(f)
	logger.ShowSQL(true)
	engine.SetLogger(logger)

	// set mapper
	tMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_db_")
	engine.SetTableMapper(tMapper)

	// sync struct to db
	//err = engine.Sync2(new(User))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// get db info
	//info, err = engine.DBMetas()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Insert
	//user := &User{
	//	Name: "magodo",
	//}
	//affected, err := engine.Insert(user)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("affected = ", affected)

	// query
	user := new(User)
	_, err = engine.Where("name = ?", "magodo").Get(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
