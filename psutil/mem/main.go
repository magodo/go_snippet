package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	mem, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mem)
}
