package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/foo", nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	req = req.WithContext(ctx)
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
	}
}
