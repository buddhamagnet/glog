package main

import (
	"github.com/buddhamagnet/glog/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/{post}", handlers.PostHandler)
    log.Println("buddhamagnet rising on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
