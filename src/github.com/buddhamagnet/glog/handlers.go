package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
	t.Execute(w, GetPosts())
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fileread, _ := ioutil.ReadFile("posts/" + r.URL.Path[1:] + ".md")
	lines := strings.Split(string(fileread), "\n")
	t := template.New("post.html")
	t, _ = t.ParseFiles("post.html")
	t.Execute(w, NewPost(lines, r.URL.Path[1:]))
}
