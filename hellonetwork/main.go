package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	help = flag.Bool("help", false, "Show this help")
)

func main() {
	flag.Parse()

	if *help || flag.NArg() != 1 {
		fmt.Println("Usage: hellonetwork [options] <address to bind to>")
		flag.PrintDefaults()
		return
	}

	addr := flag.Arg(0)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Printf("Could bind to %s: %s", addr, e)
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
	fmt.Fprintf(c, "Hi! :)\n")
}
