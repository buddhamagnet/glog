package main

import (
	"github.com/buddhamagnet/glog/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/{post}", handlers.PostHandler)
	http.ListenAndServe(":8000", r)
}
