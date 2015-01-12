package main

import (
	"github.com/buddhamagnet/glog/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/{post}", handlers.PostHandler)

	http.ListenAndServe(":8000", r)
}
