package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	port := flag.String("port", "8080", "HTTP port")
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/posts/{post}", PostHandler)
	r.HandleFunc("/pages/{page}", PageHandler)
	log.Println("buddhamagnet rising on port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
