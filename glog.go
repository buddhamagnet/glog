package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	port := flag.String("port", "8080", "HTTP port")
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/posts.rss", FeedHandler)
	r.HandleFunc("/{content}/{article}", ContentHandler)
	log.Println("buddhamagnet rising on port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
