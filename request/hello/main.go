package main

import (
	"log"

	"github.com/parnurzeal/gorequest"
)

func main() {
	resp, body, err := gorequest.New().Post("http://192.168.153.22:2245").Send(`
{
	"Action": "Bar"
}
`).End()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	log.Println(body)
}
