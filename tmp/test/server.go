package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	/* explictly bind to 0.0.0.0 to avoid hassle about ipv4 or ipv6 stuff... */
	listen, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", "12345"))
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()
	for {
		c, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			tcpConn, _ := c.(*net.TCPConn)
			fmt.Println("Close write direction.")
			tcpConn.CloseWrite()
			fmt.Println("Start reading")
			_, err := io.Copy(ioutil.Discard, tcpConn)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Finish reading")
			c.Close()
		}()
	}
}
