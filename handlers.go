package main

import (
    "fmt"
    "github.com/gorilla/feeds"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	preprocess(w)
	feed := &feeds.Feed{
	    Title: "buddhamagnet",
	    Link: &feeds.Link{Href: "http://glog.herokuapp.com"},
	    Description: "blog",
	    Author: &feeds.Author{"Dave Goodchild", "buddhamagnet@gmail.com"},
	    Created: time.Now(),
    }
	posts := GetPosts()
	feed.Items = []*feeds.Item{}
	for _, post := range posts {
		if post.Note == "excluded" {
			continue
		}
		item := &feeds.Item{
			Title: post.Title,
			Link: &feeds.Link{Href: "http://glog.herokuapp.com/" + post.File},
			Description: post.Excerpt(),
			Author: &feeds.Author{"Dave Goodchild", "buddhamagnet@gmail.com"},
			Created: time.Now(),
		}
		feed.Items = append(feed.Items, item)
	}
	rss, _ := feed.ToRss()
	fmt.Fprint(w, rss)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	preprocess(w)
	fileread, _ := ioutil.ReadFile(r.URL.Path[1:] + ".md")
	lines := strings.Split(string(fileread), "\n")
	t := template.New("post.html")
	t, _ = t.ParseFiles("post.html")
	post := NewPost(lines, r.URL.Path[1:])
	t.Execute(w, post)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	preprocess(w)
	fileread, _ := ioutil.ReadFile(r.URL.Path[1:] + ".md")
	lines := strings.Split(string(fileread), "\n")
	t := template.New("page.html")
	t, _ = t.ParseFiles("page.html")
	page := NewPage(lines, r.URL.Path[1:])
	t.Execute(w, page)
}
