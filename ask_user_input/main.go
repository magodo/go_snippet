package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)

func askUserEnter(msg string, out interface{}) {
	fmt.Print(msg)
	input, _ := r.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	switch out.(type) {
	case *string:
		*(out.(*string)) = input
	case *int:
		num, _ := strconv.Atoi(input)
		*(out.(*int)) = num
	case *uint64:
		num, _ := strconv.ParseUint(input, 10, 64)
		*(out.(*uint64)) = num
	}
}

func ChooseDBTypeID() (typeid int) {
	askUserEnter(`Choose DB type below: 
    1: 		mysql-5.5
    2: 		mysql-5.1
    6: 		mysql-5.6
    10:     mysql-5.7
    16:		postgresql-9.6
Enter index: `, &typeid)
	return
}

func ChooseHaInstanceID() (id string) {
	askUserEnter("Enter HA instance ID (len < 36): ", &id)
	return
}

func ChooseCPULimit() (limit uint64) {
	askUserEnter("Enter CPU core amount limit: ", &limit)
	return
}

func main() {
	fmt.Println(ChooseDBTypeID(), ChooseHaInstanceID(), ChooseCPULimit())
}
