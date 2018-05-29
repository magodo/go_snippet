package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/shirou/gopsutil/cpu"
)

func sum(values ...float64) (total float64) {
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	n, err := cpu.Counts(false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Counts(false): ", n)

	n, err = cpu.Counts(true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Counts(true): ", n)

	percList, err := cpu.Percent(100*time.Millisecond, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Percent per cpu: ", spew.Sdump(percList))
	fmt.Println(uint8(sum(percList...) / float64(len(percList))))
	fmt.Println(uint8(math.Ceil((sum(percList...) / float64(len(percList)) / 100.0))))

	percList, err = cpu.Percent(100*time.Millisecond, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Percent total: ", spew.Sdump(percList))

	infos, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}
	_ = infos
	//fmt.Println("Info: ", spew.Sdump(infos))

}
