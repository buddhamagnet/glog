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
	posts := GetPosts()
	t.Execute(w, posts)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fileread, _ := ioutil.ReadFile(r.URL.Path[1:] + ".md")
	lines := strings.Split(string(fileread), "\n")
	t := template.New("post.html")
	t, _ = t.ParseFiles("post.html")
	post := NewPost(lines, r.URL.Path[1:])
	t.Execute(w, post)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	fileread, _ := ioutil.ReadFile(r.URL.Path[1:] + ".md")
	lines := strings.Split(string(fileread), "\n")
	t := template.New("page.html")
	t, _ = t.ParseFiles("page.html")
	page := NewPage(lines, r.URL.Path[1:])
	t.Execute(w, page)
}
