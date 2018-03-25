package main

import (
	"fmt"
	"io"
	"log"
	"net"
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
	case 2:
		fmt.Fprintln(c1, "Oops, he/she lefts")
		c1.Close()
	}
}

func handleConn(conn io.ReadWriteCloser) {
	fmt.Fprintln(conn, "Wating for a stranger...")
	select {
	case pendingChann <- conn:
		// hand over net.Conn to another go routine
	case peer := <-pendingChann:
		chat(peer, conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("failed to create server: ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("failed to accept: ", err)
		}
		go handleConn(conn)
	}
}
