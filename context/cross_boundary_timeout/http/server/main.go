package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New request")
		if d, ok := r.Context().Deadline(); ok {
			fmt.Println(d)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
