package main

import (
	"io"
	"log"
	"net"
)

var toQuit = false

func main() {
	// TODO: flags
	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for !toQuit {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Accept: %s", conn.RemoteAddr())

		go processClient(conn)
	}
}

func processClient(conn net.Conn) {

	defer conn.Close()

	// TODO: flags
	buf := make([]byte, 5)
	is_new_msg := true

	for !toQuit {
		n, err := conn.Read(buf)

		if err != nil {
			/* `n == 0` here is for the case that non-zero count of bytes together with EOF is read,
			   should still process the data, next read is guaranteed to get 0 bytes and EOF err.*/
			if err == io.EOF && n == 0 {
				/* peer disconnect */
				log.Printf("Peer %s closed", conn.RemoteAddr())
				return
			}
			/* other error */
			log.Println("read error", err)
			return
		}

		process(buf, &is_new_msg)

		/* clear buf */
		for i := 0; i < len(buf); i++ {
			buf[i] = 0
		}
	}

	/* TODO: quit intermediately */
	postInterrupt(conn)
}

func process(buf []byte, is_new_msg *bool) {

	if *is_new_msg {
		/* TODO: Process new message, e.g. header part. Unset flag after done. */
		log.Printf("[New Message] ")
		*is_new_msg = false
	}
	log.Println(buf)
}

func postInterrupt(conn net.Conn) {
	conn.Write([]byte("Interrupt"))
}
