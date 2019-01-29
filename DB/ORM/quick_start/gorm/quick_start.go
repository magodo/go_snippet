package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"

	//"github.com/jinzhu/gorm"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string
	Created time.Time
	Updated time.Time
}

type MyInt int

type THaInstance struct {
	Id MyInt
}

type A struct {
	I int
}

type Foo struct {
	Id   int
	Name string
}

func main() {
	mysql()
}

func pg() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable")
	//db, err := gorm.Open("postgres", "host=172.255.255.254 port=5432 user=postgres dbname=postgres sslmode=disable password=123")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.LogMode(true)

	// Read
	//var loc []string
	//if err = db.Table("pg_stat_replication").Select("write_location").Scan(&loc).Error; err != nil {
	//	log.Fatal(err)
	//}

	type Foo struct {
		i int
		j int
	}
	var foo Foo
	if err = db.Raw("select i, j from b").Scan(&foo).Error; err != nil {
		log.Fatal(err)
	}

	spew.Dump(foo)
	// Delete - delete product
	//db.Delete(&product)

	//raw, _ := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable")
	//rows, _ := raw.Query("select * from a")
	//defer rows.Close()
	//for rows.Next() {
	//	var i int
	//	rows.Scan(&i)
	//	spew.Dump(i)
	//}
}

func mysql() {
	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	//var foos []Foo
	//if err = db.Table("foo").Find(&foos).Error; err != nil {
	//	log.Fatal(err)
	//}
	//spew.Dump(foos)

	type Result struct {
		Id   int
		Bar  int
		Name string
	}
	var result []Result
	if err = db.Raw("select id, name, bar from foo").Scan(&result).Error; err != nil {
		log.Fatal(err)
	}
	spew.Dump(result[1].Name)

	//var name string
	//raw, _ := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?")
	//row, _ := raw.Query("select name from foo")
	//defer row.Close()
	//for row.Next() {
	//	err := row.Scan(&name)
	//	if err != nil {
	//		spew.Dump(err)
	//	}

	//	spew.Dump(name)
	//}

}
