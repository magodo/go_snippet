package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	usage, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(usage)
}
