package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/pkg/errors"
)

//IPv4ToIPv6OfSameIterface lookup the interface having the specified ipv4, and return the binded ipv6 address.
func IPv4ToIPv6OfSameIterface(ipv4 string) (ipv6 net.IPAddr, err error) {
	itfs, err := net.Interfaces()
	if err != nil {
		err = errors.Wrap(err, "Failed to list interfaces")
		return
	}

	for _, itf := range itfs {
		var addrs []net.Addr
		addrs, err = itf.Addrs()

		if err != nil {
			err = errors.Wrapf(err, "Failed to list addresses for interface: %s", itf.Name)
			return
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.String() == ipv4 {
				/* re-loop addrs of this interface and return the ipv6 is any */
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() == nil {
						ipv6 = net.IPAddr{
							IP:   ipnet.IP,
							Zone: itf.Name,
						}
						return
					}
				}
				return
			}
		}
	}
	return
}

func main() {
	var ipv4 string
	flag.StringVar(&ipv4, "i", "", "ipv4")
	flag.Parse()
	ipv6, err := IPv4ToIPv6OfSameIterface(ipv4)
	if err != nil {
		log.Fatal(err)
	}
	port := "8776"
	fmt.Println(ipv6)
	_, err = net.Listen("tcp", net.JoinHostPort(ipv6.String(), port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second)
	}
}
