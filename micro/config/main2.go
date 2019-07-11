package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
)

const udbRegionConfigKey = "region_info"

//////
type ZoneConfig struct {
	ZoneId     int
	ZoneCode   string
	RegionId   int    `json:"az_group"`
	RegionCode string `json:"az_group_code"`
	ZKServers  []struct {
		Host string
		Port int
	} `json:"zookeeper"`
	DB struct {
		Host     string
		Port     int
		User     string `json:"super_user"`
		Password string `json:"super_password"`
		Database string
	}
	ZNodePath struct {
		UnetAccess string
	}
}

type Foo struct {
	A string
}

func main() {
	consulSource := consul.NewSource(
		consul.WithAddress("localhost:8500"),
		consul.WithPrefix("/"),
		consul.StripPrefix(true),
	)

	conf := config.NewConfig(
		config.WithSource(consulSource),
	)

	foo := []int{}
	//foo := Foo{}
	if err := conf.Get("foo").Scan(&foo); err != nil {
		log.Fatal(err)
	}
	spew.Dump(foo)

	w, err := conf.Watch("foo")
	if err != nil {
		log.Fatal(err)
	}

	v, err := w.Next()
	if err != nil {
		log.Fatal(err)
	}

	if err := v.Scan(&foo); err != nil {
		log.Fatal(err)
	}
	spew.Dump(foo)
}
