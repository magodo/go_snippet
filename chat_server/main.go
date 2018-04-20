package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
)

var pendingChann = make(chan io.ReadWriteCloser)

func chatCopy(src, dst io.ReadWriter, finishChann chan int8, id int8) {
	_, err := io.Copy(dst, src)
	if err != nil {
		finishChann <- id
	}
}

func chat(c1, c2 io.ReadWriteCloser) {
	finishChann := make(chan int8, 1) // don't block any of the two routines when finished
	bc := io.MultiWriter(c1, c2)
	fmt.Fprintln(bc, `Someone enters, say "Hello!"`)
	go chatCopy(c2, c1, finishChann, 1)
	go chatCopy(c1, c2, finishChann, 2)
	id := <-finishChann
	switch id {
	case 1:
		fmt.Fprintln(c2, "Oops, he/she lefts")
		c2.Close()
		// TODO: fix reconnect issue
		//match(c2)
	case 2:
		fmt.Fprintln(c1, "Oops, he/she lefts")
		c1.Close()
		//match(c1)
	}
}

func match(conn io.ReadWriteCloser) {
	fmt.Fprintln(conn, "Wating for a stranger...")
	select {
	case pendingChann <- conn:
		// hand over net.Conn to another go routine
	case peer := <-pendingChann:
		chat(peer, conn)
	}
}

//func main() {
//	listener, err := net.Listen("tcp", ":8000")
//	if err != nil {
//		log.Fatalln("failed to create server: ", err)
//	}
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			log.Fatalln("failed to accept: ", err)
//		}
//		go match(conn)
//	}
//}

var tpls *template.Template = template.Must(template.ParseFiles("root.html"))

type socket struct {
	io.ReadWriter
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tpls.ExecuteTemplate(w, "root.html", listenAddr)
}

func socketHandler(ws *websocket.Conn) {
	s := socket{ws, make(chan bool)}
	go match(s)
	<-s.done
}

var listenAddr = getIpAddr() + ":8000"

func getIpAddr() string {
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	fmt.Println("Hosting on: ", listenAddr)
	http.HandleFunc("/", rootHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
