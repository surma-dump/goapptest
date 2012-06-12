package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	help = flag.Bool("help", false, "Show this help")
	web  = flag.String("web", "", "Address to listen on")
)

func main() {
	flag.Parse()

	if *help || *web == "" {
		fmt.Println("Usage: hellonetwork [options] <address to bind to>")
		flag.PrintDefaults()
		return
	}

	http.HandleFunc("/", handler)
	e := http.ListenAndServe(*web, nil)
	if e != nil {
		log.Fatalf("Could bind to %s: %s", *web, e)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tach! :)\n")
}
