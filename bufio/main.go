package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf := bufio.NewReader(f)
	
	for line := range buf.ReadString(
}
