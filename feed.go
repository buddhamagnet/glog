package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/feeds"
)

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	preprocess(w)
	feed := &feeds.Feed{
		Title:       "buddhamagnet",
		Link:        &feeds.Link{Href: "http://glog.herokuapp.com"},
		Description: "blog",
		Author:      &feeds.Author{"Dave Goodchild", "buddhamagnet@gmail.com"},
		Created:     time.Now(),
	}
	posts := GetPosts()
	feed.Items = []*feeds.Item{}
	for _, post := range posts {
		if post.Note == "excluded" {
			continue
		}
		item := &feeds.Item{
			Title:       post.Title,
			Link:        &feeds.Link{Href: "http://glog.herokuapp.com/" + post.File},
			Description: post.Excerpt(),
			Author:      &feeds.Author{"Dave Goodchild", "buddhamagnet@gmail.com"},
			Created:     time.Now(),
		}
		feed.Items = append(feed.Items, item)
	}
	rss, _ := feed.ToRss()
	fmt.Fprint(w, rss)
}
