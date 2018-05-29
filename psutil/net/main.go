package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/net"
)

func main() {
	stats, err := net.Connections("all")
	if err != nil {
		log.Fatal(err)
	}
	for _, stat := range stats {
		if stat.Laddr.Port == 3306 {
			fmt.Println(stat)
		}
	}
}
