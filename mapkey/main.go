package main

import (
	"fmt"
	"foo/mapkey"
)

var map1 = map[int]int{
	1: 1,
	2: 2,
}

var map2 = map[string]string{
	"1": "1",
	"2": "2",
}

func main() {
	k, ok := mapkey.Mapkey(map1, 1)
	fmt.Println(k, ok)
	k, ok = mapkey.Mapkey(map2, "1")
	fmt.Println(k, ok)
	k, ok = mapkey.Mapkey(map2, "3")
	fmt.Println(k, ok)
}
