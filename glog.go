package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	port := flag.String("port", "8080", "HTTP port")
	flag.Parse()
	r := NewRouter()
	log.Println("buddhamagnet rising on port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
