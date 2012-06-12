package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	help   = flag.Bool("help", false, "Show this help")
	socket = flag.String("socket", "", "Socket to listen to")
)

func main() {
	flag.Parse()

	if *help || *socket == "" {
		fmt.Println("Usage: hellonetwork [options] <address to bind to>")
		flag.PrintDefaults()
		return
	}

	l, e := net.Listen("tcp", *socket)
	if e != nil {
		log.Printf("Could bind to %s: %s", *socket, e)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			continue
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Fprintf(c, "Tach! :)\n")
}
