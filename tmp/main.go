package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

func main() {
	p := regexp.MustCompile("#time (.{6})")
	matches := p.FindStringSubmatch("#time abcdef")
	fmt.Println(matches)
	t, err := time.ParseInLocation("060102 15:04:05", "060102 15:04:05", time.Local)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
}
