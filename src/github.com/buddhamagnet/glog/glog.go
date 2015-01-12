package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{post}", PostHandler)
    log.Println("buddhamagnet rising on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
