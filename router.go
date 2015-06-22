package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/", HomeHandler)
	http.Handle("/",
		GlogMiddleware(
			LoggingMiddleware(
				RecoverMiddleware(r))))
	r.HandleFunc("/posts.rss", FeedHandler)
	r.HandleFunc("/{content}/{article}", ContentHandler)
	return r
}
