/*
* var sock = new WebSocket("ws://localhost:8000/");
* sock.onmessage = function(m){ console.log("Received: ", m.data); }
* sock.send("Hello!\n")
 */

package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func handle(c *websocket.Conn) {
	var inputStr string
	fmt.Fscan(c, &inputStr)
	fmt.Println("received: ", inputStr)
	fmt.Fprint(c, "Hello")
}

func main() {
	http.Handle("/", websocket.Handler(handle))
	http.ListenAndServe(":8000", nil)
}
