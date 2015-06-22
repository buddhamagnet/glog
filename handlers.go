package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func preprocess(w http.ResponseWriter) {
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.Header().Set("X-Github", "http://github.com/buddhamagnet")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	preprocess(w)
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
	posts := GetPosts()
	t.Execute(w, posts)
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	preprocess(w)
	vars := mux.Vars(r)
	content := vars["content"]
	fileread, _ := ioutil.ReadFile(r.URL.Path[1:] + ".md")
	lines := strings.Split(string(fileread), "\n")
	t := template.New(content + ".html")
	t, _ = t.ParseFiles(content + ".html")
	post := NewContent(content, r.URL.Path[1:], lines)
	t.Execute(w, post)
}

func NewContent(content, path string, lines []string) interface{} {
	switch content {
	case "posts":
		return NewPost(lines, path)
	case "pages":
		return NewPage(lines, path)
	}
	return nil
}
